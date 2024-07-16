package products

import "github.com/labstack/echo/v4"

type Product struct {
	ID          uint
	UserID      uint
	Name        string
	Category    string
	Description string
	Price       int
	Stock       int
	ImageUrl    string
}

type Handler interface {
	AddProduct() echo.HandlerFunc
	GetProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
}

type Query interface {
	AddProduct(newProducts Product, userID uint) error
	GetProduct(ID uint) (Product, error)
	UpdateProduct(productID uint, updatedProduct Product) error
	DeleteProduct(productID uint) error
	GetAllProducts() ([]Product, error)
}

type Service interface {
	AddProduct(newProduct Product, userID uint) error
	GetProduct(ID uint) (Product, error)
	UpdateProduct(productID uint, userID uint, updatedProduct Product) error
	DeleteProduct(productID uint, userID uint) error
	GetAllProducts() ([]Product, error)
}
