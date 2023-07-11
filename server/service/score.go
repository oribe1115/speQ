package service

import (
	"context"
	"speQ/generated/model"
	"speQ/service/metrics"
)

func (s *Services) CreateNewScore(ctx context.Context, contestant string, score float64) error {
	arg := model.InsertScoreParams{
		ContestantID: contestant,
		Score:        score,
	}
	err := s.queries.InsertScore(ctx, arg)
	if err != nil {
		return err
	}

	metrics.NewScore(contestant, score)

	return nil
}
