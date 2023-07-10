package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"speQ/generated/api"
	"speQ/generated/model"
	"speQ/util"
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
	return api.ContestInfo{
		Title:              &modelValue.Title,
		Description:        util.NullString2Pointer(modelValue.Description),
		ScheduledStartTime: util.NullTime2Pointer(modelValue.ScheduledStartTime),
		StartTime:          util.NullTime2Pointer(modelValue.StartTime),
		EndTime:            util.NullTime2Pointer(modelValue.EndTime),
		VotingFreezeTime:   util.NullTime2Pointer(modelValue.VotingFreezeTime),
	}
}

func convertContestInfo2Model(apiValue *api.ContestInfo) *model.ContestInfo {
	return &model.ContestInfo{
		Title:              util.Pointer2NullString(apiValue.Title).String,
		Description:        util.Pointer2NullString(apiValue.Description),
		ScheduledStartTime: util.Pointer2NullTime(apiValue.ScheduledStartTime),
		StartTime:          util.Pointer2NullTime(apiValue.StartTime),
		EndTime:            util.Pointer2NullTime(apiValue.EndTime),
		VotingFreezeTime:   util.Pointer2NullTime(apiValue.VotingFreezeTime),
	}
}
