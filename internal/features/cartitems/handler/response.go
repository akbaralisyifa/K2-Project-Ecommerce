package handler

import "ecommerce/internal/features/products"

type ProductResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	ImageUrl    string `json:"image_url"`
}

func ToGetProductResponse(input products.Product) ProductResponse {
	return ProductResponse{
		ID:          input.ID,
		UserID:      input.UserID,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageUrl:    input.ImageUrl,
	}
}
