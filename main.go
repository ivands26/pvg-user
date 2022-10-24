package main

import (
	"fmt"
	"pvg/config"
	"pvg/factory"
	"pvg/infra/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.AddTrailingSlash())

	factory.InitFactory(e, db)
	fmt.Println("----Application is Running----")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))

}
