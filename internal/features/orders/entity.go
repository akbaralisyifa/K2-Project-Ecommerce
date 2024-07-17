package orders

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/products"
)

type Order struct {
	ID              uint
	UserID          uint
	PaymentMethod   string
	ShippingAddress string
	Status          string
	OrderItems      []OrderItems
}

type OrderItems struct {
	OrderID    uint
	ProductID  uint
	Quantity   int
	TotalPrice int
}

type Handler interface {
}

type Query interface {
	GetAllOrder(UserID uint) ([]Order, error)
	CreateOrders(newOrders Order) (uint, error)
	UpdateOrderStatus(OrderID uint, newStatus string) error
	CreateOrderItems(orderID uint, newOrderItems []cartitems.CartItem) error
	GetOrderItems(OrderID uint) ([]OrderItems, error)
	Checkout(UserID uint, newOrder Order, productID uint, updateProduct products.Product) error
}

type Service interface {
	CreateOrders(newOrders Order) error
	GetAllOrder(userID uint) ([]Order, error)
	UpdateOrderStatus(OrderID uint, newStatus string) error
	CreateOrderItems(orderID uint, newOrderItem []cartitems.CartItem) error 
	GetOrderItems(OrderID uint) ([]OrderItems, error)
}