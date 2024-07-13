package main

import (
	"ecommerce/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New();
	setup := config.InportSetting();
	connect, _ := config.ConnectDB(&setup);


	connect.Migrator()

	e.Pre(middleware.RemoveTrailingSlash());
	e.Use(middleware.Logger());
	e.Use(middleware.CORS())
}
