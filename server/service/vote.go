package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"speQ/generated/model"
	"speQ/service/metrics"
)

func (s *Services) CreateNewVote(ctx context.Context, voter string, target string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	txQueries := s.queries.WithTx(tx)
	lastVote, err := txQueries.GetLastVote(ctx, voter)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	arg := model.InsertVoteParams{
		Voter:  voter,
		Target: target,
	}

	err = txQueries.InsertVote(ctx, arg)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if lastVote.Target != "" {
		metrics.DecVoteCount(lastVote.Target)
	}
	metrics.IncVoteCount(target)

	return nil
}
