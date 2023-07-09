package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"speQ/generated/api"
)

func (r *Router) GetMe(c echo.Context) error {
	traPID, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	isRoot, err := r.services.IsRoot(ctx, traPID)
	if err != nil {
		slog.Error(err.Error())
		return echo.ErrInternalServerError
	}

	isAdmin, err := r.services.IsAdmin(ctx, traPID)
	if err != nil {
		slog.Error(err.Error())
		return echo.ErrInternalServerError
	}

	res := api.UserInfo{
		IsAdmin: isAdmin,
		IsRoot:  isRoot,
		TraPId:  traPID,
	}

	return c.JSON(http.StatusOK, res)
}

func (r *Router) PutAdminUsers(c echo.Context) error {
	_, httpErr := r.requireRootOrAdmin(c)
	if httpErr != nil {
		return httpErr
	}

	req := &api.PutAdminUsersJSONRequestBody{}
	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	newAdmins, err := r.services.RegisterAdminUsers(ctx, *req)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to update admin users: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, newAdmins)
}

func (r *Router) GetAdmins(c echo.Context) error {
	_, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	adminUsers, err := r.queries.GetAdminUsers(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get root users: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, adminUsers)
}

func (r *Router) GetRootUsers(c echo.Context) error {
	_, httpErr := requireLogin(c)
	if httpErr != nil {
		return httpErr
	}

	ctx := context.Background()
	rootUsers, err := r.queries.GetRootUsers(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get root users: %v", err))
		return echo.ErrInternalServerError
	}

	// TODO: Fix openpai schema
	return c.JSON(http.StatusOK, rootUsers[0])
}

func (r *Router) PutContestants(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) GetContestants(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
