package products

import "github.com/labstack/echo/v4"

type Product struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	ImageUrl    string `json:"image_url"`
}

type Handler interface {
	AddProduct() echo.HandlerFunc
	GetProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
	GetAllProductsByOwner() echo.HandlerFunc
}

type Query interface {
	AddProduct(newProducts Product, userID uint) error
	GetProduct(ID uint) (Product, error)
	UpdateProduct(productID uint, updatedProduct Product) error
	DeleteProduct(productID uint) error
	GetAllProducts() ([]Product, error)
	GetAllUserProducts(userID uint) ([]Product, error)
	GetAllOtherUserProducts(userID uint) ([]Product, error)
}

type Service interface {
	AddProduct(newProduct Product, userID uint) error
	GetProduct(ID uint) (Product, error)
	UpdateProduct(productID uint, userID uint, updatedProduct Product) error
	DeleteProduct(productID uint, userID uint) error
	GetAllProducts() ([]Product, error)
	GetAllUserProducts(userID uint) ([]Product, error)
	GetAllOtherUserProducts(userID uint) ([]Product, error)
}
