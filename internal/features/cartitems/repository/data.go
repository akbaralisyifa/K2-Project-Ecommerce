package repository

import (
	"ecommerce/internal/features/cartitems"

	"gorm.io/gorm"
)

type CartItems struct {
	UserID     uint
	ProductID  uint
	Qty        uint
	TotalPrice uint
}

type Products struct {
	gorm.Model
	UserID      uint
	Name        string
	Category    string
	Description string
	Price       int
	Stock       int
	ImageUrl    string
	CartItems   CartItems `gorm:"foreignKey:ProductID"`
}

type Users struct {
	gorm.Model
	Fullname   string
	Email      string
	Password   string
	Phone      string
	Address    string
	ImgProfile string
	Products   []Products  `gorm:"foreignKey:UserID"`
	CartItem   []CartItems `gorm:"foreignKey:UserID"`
}

// dari database di pindah ke entity
func (c *CartItems) ToCartItemsEntity() cartitems.CartItem {
	return cartitems.CartItem{
		UserID:     c.UserID,
		ProductID:  c.ProductID,
		Qty:        c.Qty,
		TotalPrice: c.TotalPrice,
	}
}

// dari entity pindah ke database
func ToCartItemsQuery(input cartitems.CartItem) CartItems {
	return CartItems{
		UserID:     input.UserID,
		ProductID:  input.ProductID,
		Qty:        input.Qty,
		TotalPrice: input.TotalPrice,
	}
}

// dari database di pindah ke entity
func (c *Products) ToProductsEntity() cartitems.Product {
	return cartitems.Product{
		ID:          c.ID,
		UserID:      c.UserID,
		Name:        c.Name,
		Category:    c.Category,
		Description: c.Description,
		Price:       c.Price,
		Stock:       c.Stock,
		ImageUrl:    c.ImageUrl,
	}
}

// dari entity pindah ke database
func ToProductsQuery(input cartitems.Product) Products {
	return Products{
		UserID:      input.UserID,
		Name:        input.Name,
		Category:    input.Category,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageUrl:    input.ImageUrl,
	}
}
