package handler

import "ecommerce/internal/features/orders"

type CreateOrderRequest struct {
	PaymentMethod   string
	ShippingAddress string `json:"shipping_address"`
}

type CheckoutRequest struct {
	PaymentMethod   string              `json:"payment_method"`
	ShippingAddress string              `json:"shipping_address"`
	OrderItems      []orders.OrderItems `json:"order_items"`
}

type MidTransRequest struct {
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
}

func ToModelOrders(co CreateOrderRequest, userId uint, status string) orders.Order {
	return orders.Order{
		UserID:          userId,
		PaymentMethod:   co.PaymentMethod,
		ShippingAddress: co.ShippingAddress,
		Status:          status,
	}
}

func ToModelOrders2(orderID uint, co MidTransRequest) orders.Order {
	return orders.Order{
		ID:            orderID,
		PaymentMethod: co.PaymentType,
		Status:        co.TransactionStatus,
	}
}
