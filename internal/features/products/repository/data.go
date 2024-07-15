package repository

import (
	"ecommerce/internal/features/products"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name        string
	Category    string
	Description string
	Price       int
	Stock       int
	ImageUrl    string
}

// dari database di pindah ke entity
func (p *Products) ToProductsEntity() products.Product {
	return products.Product{
		ID:          p.ID,
		Name:        p.Name,
		Category:    p.Category,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		ImageUrl:    p.ImageUrl,
	}
}

// dari entity pindah ke database
func ToProductsQuery(input products.Product) Products {
	return Products{
		Name:        input.Name,
		Category:    input.Category,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageUrl:    input.ImageUrl,
	}
}
