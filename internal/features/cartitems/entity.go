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

type User struct {
	ID         uint
	Fullname   string
	Email      string
	Password   string
	Phone      string
	Address    string
	ImgProfile string
	CartItems  []CartItem
	Products   []Product
}

type Product struct {
	ID          uint
	UserID      uint
	Name        string
	Category    string
	Description string
	Price       int
	Stock       int
	ImageUrl    string
	CartItem    CartItem
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
	UpdateCartItem(newCartItems CartItem) error
	GetProduct(productID uint) (Product, error)
}

type Service interface {
	AddCartItem(newCartItems CartItem, userID uint) error
	DeleteCartItem(productID uint, userID uint) error
	GetAllCartItems(userID uint) ([]CartItem, []Product, error)
}
