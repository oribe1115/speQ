package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
