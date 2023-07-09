package service

import (
	"context"
	"fmt"
	"speQ/generated/model"
)

func (s *Services) GetContestInfo(ctx context.Context) (*model.ContestInfo, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	txQueries := s.queries.WithTx(tx)
	rowCount, err := txQueries.CountContestInfoRow(ctx)
	if err != nil {
		return nil, err
	}

	if rowCount == 0 {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		return &model.ContestInfo{}, nil
	}

	contestInfo, err := txQueries.GetContestInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &contestInfo, nil
}
