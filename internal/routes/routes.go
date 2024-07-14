package routes

import (
	"ecommerce/internal/features/users"

	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uh users.Handler){

	c.POST("/register", uh.Register());
	c.POST("/login", uh.Login())
}