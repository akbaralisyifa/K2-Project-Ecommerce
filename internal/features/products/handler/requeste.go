package handler

import "ecommerce/internal/features/products"

type ProductInput struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	ImageUrl    string
}

func ToModelProduct(pr ProductInput) products.Product {
	return products.Product{
		Name:        pr.Name,
		Description: pr.Description,
		Category:    pr.Category,
		Price:       pr.Price,
		Stock:       pr.Stock,
		ImageUrl:    pr.ImageUrl,
	}
}
