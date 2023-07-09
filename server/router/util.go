package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func requireLogin(c echo.Context) (traPID string, httpErr *echo.HTTPError) {
	xShowcaseUser, ok := c.Request().Header["X-Showcase-User"]
	if !ok || len(xShowcaseUser) == 0 {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "You need to login")
	}

	return xShowcaseUser[0], nil
}
