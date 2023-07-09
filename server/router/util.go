package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
)

func requireLogin(c echo.Context) (traPID string, httpErr *echo.HTTPError) {
	xShowcaseUser, ok := c.Request().Header["X-Showcase-User"]
	if !ok || len(xShowcaseUser) == 0 {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "You need to login")
	}

	return xShowcaseUser[0], nil
}

func (r *Router) requireRootOrAdmin(c echo.Context) (traPID string, httpErr *echo.HTTPError) {
	traPID, httpErr = requireLogin(c)
	if httpErr != nil {
		return "", httpErr
	}

	ctx := context.Background()
	isRoot, err := r.services.IsRoot(ctx, traPID)
	if err != nil {
		slog.Error(err.Error())
		return "", echo.ErrInternalServerError
	}

	if isRoot {
		return traPID, nil
	}

	isAdmin, err := r.services.IsAdmin(ctx, traPID)
	if err != nil {
		slog.Error(err.Error())
		return "", echo.ErrInternalServerError
	}

	if isAdmin {
		return traPID, nil
	}

	return "", echo.NewHTTPError(http.StatusForbidden, "only admin or root can request")
}
