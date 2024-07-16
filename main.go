package main

import (
	"ecommerce/config"
	"ecommerce/internal/features/products"
	"ecommerce/internal/features/users"
	"ecommerce/internal/features/users/handler"
	"ecommerce/internal/features/users/repository"
	"ecommerce/internal/features/users/service"
	"ecommerce/internal/routes"
	"ecommerce/internal/utils"

	phand "ecommerce/internal/features/products/handler"
	prep "ecommerce/internal/features/products/repository"
	pserv "ecommerce/internal/features/products/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitUserRouter(db *gorm.DB) users.Handler {
	pw := utils.NewHashingPassword()
	vl := utils.NewValidatorUtility(*validator.New())
	jw := utils.NewJwtUtility()
	um := repository.NewUserModels(db)
	us := service.NewUserService(um, pw, vl, jw)
	uc := handler.NewUserController(us)

	return uc
}

func InitProductRouter(db *gorm.DB) products.Handler {
	pw := utils.NewHashingPassword()
	vl := utils.NewValidatorUtility(*validator.New())
	jw := utils.NewJwtUtility()
	pm := prep.NewProductModels(db)
	ps := pserv.NewProductService(pm, pw, vl, jw)
	pc := phand.NewProductController(ps)

	return pc
}

func main() {
	e := echo.New()
	setup := config.ImportSetting()
	connect, _ := config.ConnectDB(&setup)

	connect.AutoMigrate(&repository.Users{}, &prep.Products{})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ur := InitUserRouter(connect)
	pr := InitProductRouter(connect)
	routes.InitRoute(e, ur, pr)

	e.Logger.Fatal(e.Start(":6000"))
}
