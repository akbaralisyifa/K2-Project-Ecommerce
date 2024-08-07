package handler

import "ecommerce/internal/features/orders"

type OrdersResponse struct {
	ID              uint                `json:"id"`
	UserID          uint                `json:"user_id"`
	PaymentMethod   string              `json:"payment_method"`
	ShippingAddress string              `json:"shipping_address"`
	Status          string              `json:"status"`
	TotalOrder      uint64              `json:"total_order"`
	PaymentURL      string              `json:"payment_url"`
	OrderItems      []orders.OrderItems `json:"order_items"`
}

func ToOrderResponse(result []orders.Order) []OrdersResponse {
	toResOrder := make([]OrdersResponse, len(result))

	for i, val := range result {
		toResOrder[i] = OrdersResponse{
			ID:              val.ID,
			UserID:          val.UserID,
			PaymentMethod:   val.PaymentMethod,
			ShippingAddress: val.ShippingAddress,
			Status:          val.Status,
			TotalOrder:      val.TotalOrder,
			PaymentURL:      val.PaymentURL,
			OrderItems:      val.OrderItems,
		}
	}

	return toResOrder
}

type CheckoutResponse struct {
	PaymentURL string `json:"payment_url"`
}

func ToCheckoutResponse(input string) CheckoutResponse {
	return CheckoutResponse{
		PaymentURL: input,
	}
}
