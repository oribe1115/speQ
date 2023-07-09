package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"speQ/generated/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatal(err)
	}

	_ = model.New(db)

	e := echo.New()
	e.GET("/api/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
