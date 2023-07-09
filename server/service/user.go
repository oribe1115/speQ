package service

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"speQ/service/metrics"
	"strings"
)

func (s *Services) IsRoot(ctx context.Context, traPID string) (bool, error) {
	rootRowCount, err := s.queries.CountRowAsRoot(ctx, traPID)
	if err != nil {
		return false, fmt.Errorf("failed to get row counts as root: %v", err)
	}

	return rootRowCount > 0, nil
}

func (s *Services) IsAdmin(ctx context.Context, traPID string) (bool, error) {
	adminRowCount, err := s.queries.CountRowAsAdmin(ctx, traPID)
	if err != nil {
		return false, fmt.Errorf("failed to get row counts as admin: %v", err)
	}

	return adminRowCount > 0, nil
}

func (s *Services) RegisterRootUsers(ctx context.Context) error {
	err := s.queries.DeleteAllRootUsers(ctx)
	if err != nil {
		return err
	}
	slog.Info("dropped existing root users")

	rootUsersStr := os.Getenv("ROOT_USER")
	if rootUsersStr == "" {
		slog.Warn("ROOT_USER is empty")
		return nil
	}

	rootUsers := strings.Split(rootUsersStr, ",")

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	txQueries := s.queries.WithTx(tx)
	for _, traPID := range rootUsers {
		err = txQueries.InsertRootUser(ctx, traPID)
		if err != nil {
			return fmt.Errorf("failed to registar `%s`: %v", traPID, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("registered root users: %v", rootUsers))

	return nil
}

func (s *Services) RegisterAdminUsers(ctx context.Context, traPIDs []string) ([]string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	txQueries := s.queries.WithTx(tx)
	err = txQueries.DeleteAllAdminUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, traPID := range traPIDs {
		err = txQueries.InsertAdminUser(ctx, traPID)
		if err != nil {
			return nil, err
		}
	}

	newAdmins, err := txQueries.GetAdminUsers(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return newAdmins, nil
}

func (s *Services) RegisterContestants(ctx context.Context, traPIDs []string) ([]string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	txQueries := s.queries.WithTx(tx)
	err = txQueries.DeleteAllContestants(ctx)
	if err != nil {
		return nil, err
	}

	for _, traPID := range traPIDs {
		err = txQueries.InsertContestant(ctx, traPID)
		if err != nil {
			return nil, err
		}
	}

	newContestants, err := txQueries.GetContestants(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	metrics.ClearVoteCount(newContestants)

	return newContestants, nil
}
