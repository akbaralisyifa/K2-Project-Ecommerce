package handler

import "ecommerce/internal/features/cartitems"

type AllCartItemResponse struct {
	BuyerID        uint   `json:"buyer_id"`
	ProductOwnerID uint   `json:"product_owner_id"`
	ProductID      uint   `json:"product_id"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	Stock          int    `json:"stock"`
	ImageUrl       string `json:"image_url"`
	Qty            uint   `json:"qty"`
	TotalPrice     uint   `json:"total_price"`
}

func ToGetAllCartItemsResponse(inputp cartitems.Product, inputc cartitems.CartItem) AllCartItemResponse {
	return AllCartItemResponse{
		BuyerID:        inputc.UserID,
		ProductOwnerID: inputp.UserID,
		ProductID:      inputp.ID,
		Name:           inputp.Name,
		Category:       inputp.Category,
		Description:    inputp.Description,
		Price:          inputp.Price,
		Stock:          inputp.Stock,
		ImageUrl:       inputp.ImageUrl,
		Qty:            inputc.Qty,
		TotalPrice:     inputc.TotalPrice,
	}
}
