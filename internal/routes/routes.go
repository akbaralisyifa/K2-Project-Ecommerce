package routes

import (
	"ecommerce/config"
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/features/products"
	"ecommerce/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uh users.Handler, ph products.Handler, ch cartitems.Handler, oh orders.Handler) {
	// jwt key
	secrateJwt := config.ImportSetting().JWTSecrat
	c.POST("/register", uh.Register())
	c.POST("/login", uh.Login())
	c.GET("/products/:id", ph.GetProduct())
	c.GET("/products", ph.GetAllProducts())

	//user route
	ug := c.Group("/users")
	ug.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	ug.GET("", uh.GetUser())
	ug.PUT("", uh.UpdateUser())
	ug.DELETE("", uh.DeleteUser())

	//product route
	pg := c.Group("/products")
	pg.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	pg.GET("/productbyuser", ph.GetAllProductsByOwner())
	pg.POST("", ph.AddProduct())
	pg.PUT("/:id", ph.UpdateProduct())
	pg.DELETE("/:id", ph.DeleteProduct())

	//cart item route
	cg := c.Group("/cartitems")
	cg.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	cg.POST("", ch.AddCartItem())
	cg.GET("", ch.GetAllCartItems())
	cg.DELETE("/:id", ch.DeleteCartItem())

	// checkout
	c.POST("/checkout", oh.Checkout(), echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	c.GET("/orders", oh.GetAllOrder(), echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	c.GET("/order-history", oh.GetAllOrderHistory(), echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
}
