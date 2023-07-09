package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	e.GET("/api/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
