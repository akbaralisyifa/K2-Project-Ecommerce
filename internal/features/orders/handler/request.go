package handler

import "ecommerce/internal/features/orders"

type CreateOrderRequest struct{
	PaymentMethod 	string 	`json:"payment_method"`
	ShippingAddress string `json:"shipping_address"`
};

type CheckoutRequest struct {
	PaymentMethod   string            `json:"payment_method"`
	ShippingAddress string            `json:"shipping_address"`
	OrderItems      []orders.OrderItems `json:"order_items"`
}

func ToModelOrders(co CreateOrderRequest, userId uint, status string) orders.Order{
	return orders.Order{
		UserID: userId,
		PaymentMethod: co.PaymentMethod,
		ShippingAddress: co.ShippingAddress,
		Status: status,
	}
}