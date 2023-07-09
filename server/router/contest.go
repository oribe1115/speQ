package router

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"speQ/generated/api"
	"speQ/generated/model"
)

func (r *Router) PutContestInfo(c echo.Context) error {
	_, httpErr := r.requireRootOrAdmin(c)
	if httpErr != nil {
		return httpErr
	}

	// TODO: Validate if the contest is running

	req := &api.PutContestInfoJSONRequestBody{}
	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	newContestInfo, err := r.services.UpdateContestInfo(ctx, convertContestInfo2Model(req))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to update contest info: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, convertContestInfo2api(newContestInfo))
}

func (r *Router) GetContestInfo(c echo.Context) error {
	ctx := context.Background()

	contestInfo, err := r.services.GetContestInfo(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get contest info: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, convertContestInfo2api(contestInfo))
}

func convertContestInfo2api(modelValue *model.ContestInfo) api.ContestInfo {
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

func convertContestInfo2Model(apiValue *api.ContestInfo) *model.ContestInfo {
	modelValue := &model.ContestInfo{}

	if apiValue.Title != nil {
		modelValue.Title = *apiValue.Title
	}
	if apiValue.Description != nil {
		modelValue.Description = sql.NullString{
			String: *apiValue.Description,
			Valid:  true,
		}
	}
	if apiValue.ScheduledStartTime != nil {
		modelValue.ScheduledStartTime = sql.NullTime{
			Time:  *apiValue.ScheduledStartTime,
			Valid: true,
		}
	}
	if apiValue.StartTime != nil {
		modelValue.StartTime = sql.NullTime{
			Time:  *apiValue.StartTime,
			Valid: true,
		}
	}
	if apiValue.EndTime != nil {
		modelValue.EndTime = sql.NullTime{
			Time:  *apiValue.EndTime,
			Valid: true,
		}
	}
	if apiValue.VotingFreezeTime != nil {
		modelValue.VotingFreezeTime = sql.NullTime{
			Time:  *apiValue.VotingFreezeTime,
			Valid: true,
		}
	}

	return modelValue
}
