package repository

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID          uint
	PaymentMethod   string
	ShippingAddress string
	Status          string
	TotalOrder      uint64
	PaymentURL      string
	OrderItems      []OrderItems `gorm:"foreignKey:order_id"`
}

type OrderItems struct {
	OrderID    uint
	ProductID  uint
	Quantity   int
	TotalPrice int
}

func (or *Orders) ToOrderEntity() orders.Order {
	return orders.Order{
		ID:              or.ID,
		UserID:          or.UserID,
		PaymentMethod:   or.PaymentMethod,
		ShippingAddress: or.ShippingAddress,
		Status:          or.Status,
		TotalOrder:      or.TotalOrder,
		PaymentURL:      or.PaymentURL,
		OrderItems:      nil,
	}
}

func ToOrderQuery(input orders.Order) Orders {
	return Orders{
		UserID:          input.UserID,
		PaymentMethod:   input.PaymentMethod,
		ShippingAddress: input.ShippingAddress,
		Status:          input.Status,
		PaymentURL:      input.PaymentURL,
		TotalOrder:      input.TotalOrder,
	}
}

func (oi *OrderItems) ToOrderItemEntity() orders.OrderItems {
	return orders.OrderItems{
		OrderID:    oi.OrderID,
		ProductID:  oi.ProductID,
		Quantity:   oi.Quantity,
		TotalPrice: oi.TotalPrice,
	}
}

func ToOrderItemQuery(input orders.OrderItems) OrderItems {
	return OrderItems{
		OrderID:    input.OrderID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		TotalPrice: input.TotalPrice,
	}
}

func ToOrderItemsQuery(orderID uint, inputOrders []cartitems.CartItem) []OrderItems {
	ordItems := make([]OrderItems, len(inputOrders))

	for i, val := range inputOrders {
		ordItems[i] = OrderItems{
			OrderID:    orderID,
			ProductID:  val.ProductID,
			Quantity:   int(val.Qty),
			TotalPrice: int(val.TotalPrice),
		}
	}

	return ordItems
}

// function get All Order Items
func ToOrderItemsGetAll(orderItemList []OrderItems) []orders.OrderItems {
	orderItemsEntity := make([]orders.OrderItems, len(orderItemList))

	for i, val := range orderItemList {
		orderItemsEntity[i] = val.ToOrderItemEntity()
	}

	return orderItemsEntity
}

// function get all order
func ToOrderEntityGetAll(orderList []Orders) []orders.Order {
	orderEntity := make([]orders.Order, len(orderList))

	for i, val := range orderList {
		orderEntity[i] = val.ToOrderListGetAll()
	}

	return orderEntity
}

// function get all order list
func (or *Orders) ToOrderListGetAll() orders.Order {
	allOrderItems := or.ToOrderEntity()

	if len(or.OrderItems) > 0 {
		allOrderItems.OrderItems = make([]orders.OrderItems, len(or.OrderItems))
		for i, val := range or.OrderItems {
			allOrderItems.OrderItems[i] = orders.OrderItems{
				OrderID:    val.OrderID,
				ProductID:  val.ProductID,
				Quantity:   val.Quantity,
				TotalPrice: val.TotalPrice,
			}
		}
		allOrderItems.TotalOrder = countTotalPriceOrder(or.OrderItems)
	}

	return allOrderItems
}

func countTotalPriceOrder(orderItems []OrderItems) uint64 {
	var total uint64 = 0
	for _, item := range orderItems {
		total += uint64(item.TotalPrice)
	}
	return total
}
