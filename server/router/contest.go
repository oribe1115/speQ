package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"speQ/generated/api"
	"speQ/generated/model"
)

func (r *Router) PutContestInfo(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) GetContestInfo(c echo.Context) error {
	ctx := context.Background()

	contestInfo, err := r.services.GetContestInfo(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get contest info: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, convertContestInfo(contestInfo))
}

func convertContestInfo(modelValue *model.ContestInfo) api.ContestInfo {
	apiValue := api.ContestInfo{}
	apiValue.Title = &modelValue.Title

	if modelValue.Description.Valid {
		apiValue.Description = &modelValue.Description.String
	}
	if modelValue.ScheduledStartTime.Valid {
		apiValue.ScheduledStartTime = &modelValue.ScheduledStartTime.Time
	}
	if modelValue.StartTime.Valid {
		apiValue.StartTime = &modelValue.StartTime.Time
	}
	if modelValue.EndTime.Valid {
		apiValue.EndTime = &modelValue.EndTime.Time
	}
	if modelValue.VotingFreezeTime.Valid {
		apiValue.VotingFreezeTime = &modelValue.VotingFreezeTime.Time
	}

	return apiValue
}
