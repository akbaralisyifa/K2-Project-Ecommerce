package handler

import "ecommerce/internal/features/products"

type ProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func ToModelProduct(pr ProductInput) products.Product {
	return products.Product{
		Name:        pr.Name,
		Description: pr.Description,
		Category:    pr.Category,
		Price:       pr.Price,
		Stock:       pr.Stock,
	}
}
