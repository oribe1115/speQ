package router

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (r *Router) GetAdmins(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) GetRootUsers(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) PutContestants(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *Router) GetContestants(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
