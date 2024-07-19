package orders

import (
	"ecommerce/internal/features/cartitems"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID              uint         `json:"id"`
	UserID          uint         `json:"user_id"`
	PaymentMethod   string       `json:"payment_method"`
	ShippingAddress string       `json:"shipping_address"`
	Status          string       `json:"status"`
	TotalOrder      uint64       `json:"total_order"`
	OrderItems      []OrderItems `json:"order_items"`
	PaymentURL      string       `json:"payment_url"`
}

type OrderItems struct {
	OrderID    uint `json:"order_id"`
	ProductID  uint `json:"product_id"`
	Quantity   int  `json:"quantity"`
	TotalPrice int  `json:"total_price"`
}

type Handler interface {
	GetAllOrder() echo.HandlerFunc
	GetAllOrderHistory() echo.HandlerFunc
	Checkout() echo.HandlerFunc
	UpdateOrder() echo.HandlerFunc
}

type Query interface {
	GetAllOrder(UserID uint) ([]Order, error)
	CreateOrders(newOrders Order) (uint, error)
	UpdateOrder(OrderID uint, updateOrder Order) error
	CreateOrderItems(orderID uint, newOrderItems []cartitems.CartItem) error
	GetOrderItems(OrderID uint) ([]OrderItems, error)
	Checkout(UserID uint, newOrder Order, cartItems []cartitems.CartItem) (uint, error)
	GetAllOrderHistory(userID uint, orderID uint) ([]Order, error)
	GetOrder(OrderID uint) (Order, error)
	GetTotalOrderPrice(orderID uint) (int, error)
}

type Service interface {
	CreateOrders(newOrders Order) error
	GetAllOrder(userID uint) ([]Order, error)
	UpdateOrder(OrderID uint, updateOrder Order) error
	CreateOrderItems(orderID uint, newOrderItem []cartitems.CartItem) error
	GetOrderItems(OrderID uint) ([]OrderItems, error)
	Checkout(UserID uint, newOrders Order, artItems []cartitems.CartItem) (string, error)
	GetAllOrderHistory(userID uint, orderID uint) ([]Order, error)
}
