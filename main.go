package main

import (
	"ecommerce/config"
	"ecommerce/internal/features/users"
	"ecommerce/internal/features/users/handler"
	"ecommerce/internal/features/users/repository"
	"ecommerce/internal/features/users/service"
	"ecommerce/internal/routes"
	"ecommerce/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
);


func InitUserRouter(db *gorm.DB) users.Handler{
	pw := utils.NewHashingPassword();
	vl := utils.NewValidatorUtility(*validator.New());
	jw := utils.NewJwtUtility();
	um := repository.NewUserModels(db);
	us := service.NewUserService(um, pw, vl, jw);
	uc := handler.NewUserController(us);

	return uc;
}

func main() {
	e := echo.New();
	setup := config.InportSetting();
	connect, _ := config.ConnectDB(&setup);


	connect.AutoMigrate(&repository.Users{});

	e.Pre(middleware.RemoveTrailingSlash());
	e.Use(middleware.Logger());
	e.Use(middleware.CORS());
	
	ur := InitUserRouter(connect);

	routes.InitRoute(e, ur)

	e.Logger.Fatal(e.Start("5000"))
}
