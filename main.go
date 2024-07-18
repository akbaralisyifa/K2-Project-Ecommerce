package main

import (
	"ecommerce/config"
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
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

	chand "ecommerce/internal/features/cartitems/handler"
	crep "ecommerce/internal/features/cartitems/repository"
	cserv "ecommerce/internal/features/cartitems/service"

	oHand "ecommerce/internal/features/orders/handler"
	oQry "ecommerce/internal/features/orders/repository"
	oSrv "ecommerce/internal/features/orders/service"

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
};

func InitCartItemRouter(db *gorm.DB) cartitems.Handler {
	cm := crep.NewCartItemModels(db)
	cs := cserv.NewCartItemService(cm)
	cc := chand.NewCartItemController(cs)

	return cc
}

func InitOrderRouter(db *gorm.DB) orders.Handler {
	cm := crep.NewCartItemModels(db)
	cs := cserv.NewCartItemService(cm)
	pm := prep.NewProductModels(db);
	om := oQry.NewOrderModels(db, cm, pm)
	os := oSrv.NewOrderService(om)
	oc := oHand.NewOrderController(os, cs)

	return oc;
}

func main() {
	e := echo.New()
	setup := config.ImportSetting()
	connect, _ := config.ConnectDB(&setup)

	connect.AutoMigrate(&repository.Users{}, &prep.Products{}, &crep.CartItems{}, &oQry.Orders{}, &oQry.OrderItems{})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ur := InitUserRouter(connect)
	pr := InitProductRouter(connect)
	cr := InitCartItemRouter(connect)
	or := InitOrderRouter(connect)
	routes.InitRoute(e, ur, pr, cr, or)

	e.Logger.Fatal(e.Start(":6000"))
}