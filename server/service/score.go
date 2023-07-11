package service

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
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

func (s *Services) SetupScoreMetrics(ctx context.Context) error {
	contestants, err := s.queries.GetContestants(ctx)
	if err != nil {
		return err
	}

	scorePerContestant := map[string]float64{}
	for _, contestant := range contestants {
		scorePerContestant[contestant] = 0
	}

	latestScores, err := s.queries.GetLatestScores(ctx)
	if err != nil {
		return err
	}

	for _, scoreData := range latestScores {
		scorePerContestant[scoreData.ContestantID] = scoreData.Score
	}

	metrics.SetScore(scorePerContestant)

	slog.Info(fmt.Sprintf("setup current score: %v", scorePerContestant))

	return nil
}
