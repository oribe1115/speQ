package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"speQ/generated/api"
	"speQ/generated/model"
	"speQ/router"
	"speQ/service"
	"speQ/util"
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
		User:                 util.ReadEnvs("NS_MARIADB_USER", "DB_USERNAME"),
		Passwd:               util.ReadEnvs("NS_MARIADB_PASSWORD", "DB_PASSWORD"),
		Addr:                 fmt.Sprintf("%s:%s", util.ReadEnvs("NS_MARIADB_HOSTNAME", "DB_HOSTNAME"), util.ReadEnvs("NS_MARIADB_PORT", "DB_PORT")),
		DBName:               util.ReadEnvs("NS_MARIADB_DATABASE", "DB_DATABASE"),
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
