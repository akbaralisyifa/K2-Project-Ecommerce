package cartitems

import (
	"github.com/labstack/echo/v4"
)

type CartItem struct {
	UserID     uint
	ProductID  uint
	Qty        uint
	TotalPrice uint
}

type Handler interface {
	AddCartItem() echo.HandlerFunc
	DeleteCartItem() echo.HandlerFunc
	GetAllCartItems() echo.HandlerFunc
}

type Query interface {
	GetCartItem(productID uint, userID uint) (CartItem, error)
	AddCartItem(newCartItems CartItem) error
	GetAllCartItems(userID uint) ([]CartItem, error)
	DeleteCartItem(productID uint, userID uint) error
	//	GetProduct(productID uint) (pent.Product, error)
	UpdateCartItem(newCartItems CartItem) error
}

type Service interface {
	AddCartItem(newCartItems CartItem, userID uint) error
	DeleteCartItem(productID uint, userID uint) error
	GetAllCartItems(userID uint) ([]CartItem, error)
}
