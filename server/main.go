package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"speQ/generated/api"
	"speQ/generated/model"
	"speQ/router"
	"speQ/service"

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

	queries := model.New(db)
	services := service.NewService(db, queries)
	routerInstance := router.NewRouter(queries, services)

	ctx := context.Background()
	err = services.RegisterRootUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/api/metrics", echo.WrapHandler(promhttp.Handler()))

	apiInstance := e.Group("/api")
	api.RegisterHandlers(apiInstance, routerInstance)

	e.Logger.Fatal(e.Start(":3000"))
}
