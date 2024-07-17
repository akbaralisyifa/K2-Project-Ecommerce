package orders

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
	CreateOrders(newOrders Order) error
	UpdateOrderStatus(OrderID uint, newStatus string) error
	CreateOrderItems(newOrderItems OrderItems) error
	GetOrderItems(OrderID uint) ([]OrderItems, error)
}

type Service interface {
	CreateOrders(newOrders Order) error
	GetAllOrder(userID uint) ([]Order, error)
	UpdateOrderStatus(OrderID uint, newStatus string) error
	CreateOrderItems(newOrderItem OrderItems) error
	GetOrderItems(OrderID uint) ([]OrderItems, error)
}