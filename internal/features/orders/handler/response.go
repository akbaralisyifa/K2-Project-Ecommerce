package handler

import "ecommerce/internal/features/orders"

type OrdersResponse struct {
	ID          	uint   `json:"id"`
	UserID      	uint   `json:"user_id"`
	PaymentMethod 	string `json:"payment_method"`
	ShippingAddress string `json:"shipping_address"`
	Status       	string `json:"status"`
	OrderItems      []orders.OrderItems `json:"order_items"`
};

// func ToOrderResponse(result orders.Order) OrdersResponse{
// 	return OrdersResponse{
// 		ID: result.ID,
// 		UserID: result.UserID,
// 		PaymentMethod: result.PaymentMethod,
// 		ShippingAddress: result.ShippingAddress,
// 		Status: result.Status,
// 		OrderItems: result.OrderItems,
// 	}
// }