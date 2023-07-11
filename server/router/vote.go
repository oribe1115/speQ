package router

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"speQ/generated/api"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

func (r *Router) GetVote(c echo.Context) error {
	traPID, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	latestVote, err := r.queries.GetLatestVoteByVoter(ctx, traPID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusOK, "")
		}
		slog.Error(fmt.Sprintf("failed to GetLatestVoteByVoter: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, latestVote.Target)
}

func (r *Router) PostVote(c echo.Context) error {
	// TODO: Improve logic

	traPID, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	var req api.PostVoteJSONRequestBody
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	contestants, err := r.queries.GetContestants(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get contestants: %v", err))
		return echo.ErrInternalServerError
	}

	if !slices.Contains(contestants, req) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`%s` is not contestant", req))
	}

	err = r.services.CreateNewVote(ctx, traPID, req)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create new vote: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, req)
}

func (r *Router) GetVoteStats(c echo.Context) error {
	_, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	latestVotes, err := r.queries.GetLatestVotes(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to GetLatestVotes: %v", err))
		return echo.ErrInternalServerError
	}

	votersMap := map[string][]string{}
	for _, vote := range latestVotes {
		votersMap[vote.Target] = append(votersMap[vote.Target], vote.Voter)
	}

	res := make([]api.VoteStatsItem, 0)
	contestants, err := r.queries.GetContestants(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to GetContestants(: %v", err))
		return echo.ErrInternalServerError
	}

	for _, contestant := range contestants {
		voters, ok := votersMap[contestant]
		if !ok {
			voters = []string{}
		}
		item := api.VoteStatsItem{
			Contestant: contestant,
			Voters:     voters,
		}
		res = append(res, item)
	}

	return c.JSON(http.StatusOK, res)
}

func (r *Router) GetVoteTriple(c echo.Context) error {
	traPID, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	tripleVotes, err := r.queries.GetTripleVotesByVoter(ctx, traPID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to GetTripleVotesByVoter: %v", err))
		return echo.ErrInternalServerError
	}

	res := api.TripleVote{}
	for _, vote := range tripleVotes {
		switch vote.Order {
		case 1:
			res.First = vote.Target
		case 2:
			res.Second = vote.Target
		case 3:
			res.Third = vote.Target
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (r *Router) PostVoteTriple(c echo.Context) error {
	// TODO: Improve logic

	traPID, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	var req api.PostVoteTripleJSONRequestBody
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	contestants, err := r.queries.GetContestants(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get contestants: %v", err))
		return echo.ErrInternalServerError
	}

	if !slices.Contains(contestants, req.First) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`%s` is not contestant", req.First))
	}
	if !slices.Contains(contestants, req.Second) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`%s` is not contestant", req.Second))
	}
	if !slices.Contains(contestants, req.Third) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`%s` is not contestant", req.Third))
	}

	orderToTarget := map[int]string{
		1: req.First,
		2: req.Second,
		3: req.Third,
	}

	err = r.services.CreateNewTripleVote(ctx, traPID, orderToTarget)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create new triple vote: %v", err))
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}
