package service

import (
	"context"
	"fmt"
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
