package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"speQ/generated/api"
	"speQ/generated/model"
	"speQ/router"
	"speQ/service"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_DATABASE"),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
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
