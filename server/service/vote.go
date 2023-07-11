package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"
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

func (s *Services) SetupVoteCountMetrics(ctx context.Context) error {
	contestants, err := s.queries.GetContestants(ctx)
	if err != nil {
		return err
	}

	votesPerContestant := map[string]int{}
	for _, contestant := range contestants {
		votesPerContestant[contestant] = 0
	}

	latestVotes, err := s.queries.GetLatestVotes(ctx)
	if err != nil {
		return err
	}

	for _, vote := range latestVotes {
		votesPerContestant[vote.Target]++
	}

	metrics.SetVoteCounts(votesPerContestant)

	slog.Info(fmt.Sprintf("setup current vote counts %v", votesPerContestant))

	return nil
}
