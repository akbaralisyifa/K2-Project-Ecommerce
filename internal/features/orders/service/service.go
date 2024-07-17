package service

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/features/products"
	"errors"
	"log"
)

type OrderService struct {
	qry orders.Query
};


func NewOrderService(q orders.Query) orders.Service {
	return &OrderService{
		qry: q,
	}
};

func (os *OrderService) CreateOrders(newOrders orders.Order) error {
	_, err := os.qry.CreateOrders(newOrders)
	if err != nil {
		log.Println("create orders sql error:", err.Error())
		return  errors.New("internal server error")
	}

	return nil;
};

func (os *OrderService) GetAllOrder(userID uint) ([]orders.Order, error) {
	result, err := os.qry.GetAllOrder(userID);
	if err != nil {
		log.Println("get orders sql error:", err.Error())
		return []orders.Order{}, errors.New("internal server error")
	}

	return result, nil
};

func (os *OrderService) UpdateOrderStatus(OrderID uint, newStatus string) error {
	err := os.qry.UpdateOrderStatus(OrderID, newStatus);
	if err != nil {
		log.Println("update orders sql error:", err.Error())
		return errors.New("internal server error")
	};

	return nil;
};

func (os *OrderService)  CreateOrderItems(orderID uint, newOrderItem []cartitems.CartItem) error {
	err := os.qry.CreateOrderItems(orderID, newOrderItem);
	if err != nil {
		log.Println("create order items sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil;
}

func (os *OrderService) GetOrderItems(OrderID uint) ([]orders.OrderItems, error){
	result, err := os.qry.GetOrderItems(OrderID);
	if err != nil {
		log.Println("Get order items sql error:", err.Error())
		return []orders.OrderItems{}, errors.New("internal server error")
	};

	return result, err;
};

func (os *OrderService) Checkout(UserID uint, newOrders orders.Order, ProductID uint, updateProduct products.Product) error {

	err := os.qry.Checkout(UserID, newOrders, ProductID, updateProduct);

	if err != nil {
		log.Println("Checkout orders error", err.Error())
		return errors.New("interal server error")
	};

	return nil;
}
