package repository

import "ecommerce/internal/features/cartitems"

type CartItems struct {
	UserID     uint
	ProductID  uint
	Qty        uint
	TotalPrice uint
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
