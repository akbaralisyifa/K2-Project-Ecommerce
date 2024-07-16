package handler

import "ecommerce/internal/features/products"

type ProductResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"fullname"`
	Category    string `json:"email"`
	Price       int    `json:"phone"`
	Stock       int    `json:"address"`
}

func ToGetProductResponse(input products.Product) ProductResponse {
	return ProductResponse{
		ID:          input.ID,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Stock:       input.Stock,
	}
}
