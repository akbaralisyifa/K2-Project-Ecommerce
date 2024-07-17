package handler

import carts "ecommerce/internal/features/cartitems"

type CartItemInput struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Qty       uint `json:"qty" form:"qty"`
}

func ToModelCartItem(pr CartItemInput) carts.CartItem {
	return carts.CartItem{
		ProductID: pr.ProductID,
		Qty:       pr.Qty,
	}
}
