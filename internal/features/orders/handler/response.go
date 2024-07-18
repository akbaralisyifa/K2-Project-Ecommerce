package handler

import "ecommerce/internal/features/orders"

type OrdersResponse struct {
	ID          	uint   `json:"id"`
	UserID      	uint   `json:"user_id"`
	PaymentMethod 	string `json:"payment_method"`
	ShippingAddress string `json:"shipping_address"`
	Status       	string `json:"status"`
	TotalOrder		uint64 `json:"total_order"`
	OrderItems      []orders.OrderItems `json:"order_items"`
};

func ToOrderResponse(result []orders.Order) []OrdersResponse{
	toResOrder := make([]OrdersResponse, len(result));

	for i, val := range result{
		toResOrder[i] = OrdersResponse{
			ID: 			 val.ID,
			UserID: 		 val.UserID,
			PaymentMethod:   val.PaymentMethod,
			ShippingAddress: val.ShippingAddress,
			Status: 		 val.Status,
			TotalOrder: 	 val.TotalOrder,	
			OrderItems:		 val.OrderItems,
		}
	}

	return toResOrder;
}