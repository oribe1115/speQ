package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
	"net/http"
	"speQ/generated/api"
)

func (r *Router) PostScore(c echo.Context) error {
	// TODO: Improve logic

	_, httpErr := r.requireRootOrAdmin(c)
	if httpErr != nil {
		return httpErr
	}

	var req api.PostScoreJSONRequestBody
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

	if !slices.Contains(contestants, req.ContestantId) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`%s` is not contestant", req.ContestantId))
	}

	err = r.services.CreateNewScore(ctx, req.ContestantId, float64(req.Score))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create new score: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, req)
}
