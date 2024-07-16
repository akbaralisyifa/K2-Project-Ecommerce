package repository

import (
	"ecommerce/internal/features/orders"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID          uint
	PaymentMethod   string
	ShippingAddress string
	Status			string
	OrderItems []OrderItems `gorm:"foreignKey:order_id"`
};

type OrderItems struct {
	OrderID	   uint
	ProductID  uint
	Quantity   int
	TotalPrice int
}

func (or *Orders) ToOrderEntity() orders.Order {
	return orders.Order{
		ID: 			 or.UserID,
		UserID: 		 or.UserID,
		PaymentMethod: 	 or.PaymentMethod,
		ShippingAddress: or.ShippingAddress,
		OrderItems: 	 nil,
	}
}

func ToOrderQuery(input orders.Order) Orders {
	return Orders{
		UserID: input.UserID,
		PaymentMethod: input.PaymentMethod,
		ShippingAddress: input.ShippingAddress,
	}
}

func (oi *OrderItems) ToOrderItemEntity() orders.OrderItems {
	return orders.OrderItems{
		OrderID: 	oi.OrderID,
		ProductID:  oi.ProductID,
		Quantity:   oi.Quantity,
		TotalPrice: oi.TotalPrice,
	}
}

func ToOrderItemsQuery(input orders.OrderItems) OrderItems{
	return OrderItems{
		OrderID: input.OrderID,
		ProductID: input.ProductID,
		Quantity: input.Quantity,
		TotalPrice: input.TotalPrice,
	}
}

// function get All Order Items
func ToOrderItemsGetAll(orderItemList []OrderItems) []orders.OrderItems{
	orderItemsEntity := make([]orders.OrderItems, len(orderItemList));

	for i, val := range orderItemList{
		orderItemsEntity[i] = val.ToOrderItemEntity()
	}

	return orderItemsEntity;
}


// function get all order 
func ToOrderEntityGetAll(orderList []Orders) []orders.Order {
	orderEntity := make([]orders.Order, len(orderList));

	for i, val := range orderList{
		orderEntity[i] = val.ToOrderListGetAll()
	}

	return orderEntity;
};

// function get all order list
func (or *Orders) ToOrderListGetAll() orders.Order {
	allOrderItems := or.ToOrderEntity();

	if len(or.OrderItems) > 0 {
		allOrderItems.OrderItems = make([]orders.OrderItems, len(or.OrderItems))
		for i, val := range or.OrderItems {
			allOrderItems.OrderItems[i] = orders.OrderItems{
				OrderID: val.OrderID,
				ProductID: val.ProductID,
				Quantity: val.Quantity,
				TotalPrice: val.TotalPrice,
			}
		}
	}

	return allOrderItems
}

// func (a *Articles) ToArticlesEntityComments() articles.Article {
// 	articlesEntity := a.ToArticlesEntity()

// 	if len(a.Comments) > 0 {
// 		articlesEntity.Comments = make([]articles.Comment, len(a.Comments))
// 		for i, val := range a.Comments {
// 			articlesEntity.Comments[i] = articles.Comment{
// 				UserID:  val.UserID,
// 				Comment: val.Comment,
// 			}
// 		}
// 	}

// 	return articlesEntity
// }


// func ToArticlesEntityGetAll(articlesList []Articles) []articles.Article {
// 	articlesEntity := make([]articles.Article, len(articlesList))

// 	for i, val := range articlesList {
// 		articlesEntity[i] = val.ToArticlesEntityComments()
// 	}

// 	return articlesEntity
// }
