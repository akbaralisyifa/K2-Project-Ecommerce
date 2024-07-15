package routes

import (
	"ecommerce/config"
	"ecommerce/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uh users.Handler){
	// jwt key
	secrateJwt := config.ImportSetting().JWTSecrat
	c.POST("/register", uh.Register());
	c.POST("/login", uh.Login());

	ug := c.Group("/users");
	ug.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	));

	ug.GET("", uh.GetUser());
	ug.PUT("", uh.UpdateUser());
	ug.DELETE("", uh.DeleteUser());
}