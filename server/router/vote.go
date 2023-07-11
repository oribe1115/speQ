package router

import (
	"context"
	"fmt"
	"net/http"
	"speQ/generated/api"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

func (r *Router) GetVote(c echo.Context) error {
	//TODO implement me
	return echo.ErrNotImplemented
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
	//TODO implement me
	return echo.ErrNotImplemented
}

func (r *Router) GetVoteTriple(c echo.Context) error {
	//TODO implement me
	return echo.ErrNotImplemented
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
