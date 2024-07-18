package service

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/utils"
	"errors"
	"log"
)

type OrderService struct {
	qry orders.Query
}

func NewOrderService(q orders.Query) orders.Service {
	return &OrderService{
		qry: q,
	}
}

func (os *OrderService) CreateOrders(newOrders orders.Order) error {
	_, err := os.qry.CreateOrders(newOrders)
	if err != nil {
		log.Println("create orders sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (os *OrderService) GetAllOrder(userID uint) ([]orders.Order, error) {
	result, err := os.qry.GetAllOrder(userID)
	if err != nil {
		log.Println("get orders sql error:", err.Error())
		return []orders.Order{}, errors.New("internal server error")
	}

	return result, nil
}

func (os *OrderService) GetAllOrderHistory(userID uint) ([]orders.Order, error) {

	result, err := os.qry.GetAllOrderHistory(userID)
	if err != nil {
		log.Println("get order history error:", err.Error())
		return []orders.Order{}, errors.New("internal erver error")
	}

	return result, nil
}

// UpdateOrder(OrderID uint, updateOrder orders.Order) error
func (os *OrderService) UpdateOrder(OrderID uint, updateOrders orders.Order) error {
	err := os.qry.UpdateOrder(OrderID, updateOrders)
	if err != nil {
		log.Println("update orders sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (os *OrderService) CreateOrderItems(orderID uint, newOrderItem []cartitems.CartItem) error {
	err := os.qry.CreateOrderItems(orderID, newOrderItem)
	if err != nil {
		log.Println("create order items sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (os *OrderService) GetOrderItems(OrderID uint) ([]orders.OrderItems, error) {
	result, err := os.qry.GetOrderItems(OrderID)
	if err != nil {
		log.Println("Get order items sql error:", err.Error())
		return []orders.OrderItems{}, errors.New("internal server error")
	}

	return result, err
}

// Checkout(UserID uint, newOrder orders.Order, productID uint, cartItems []cartitems.CartItem ) error
func (os *OrderService) Checkout(UserID uint, newOrder orders.Order, cartItems []cartitems.CartItem) (string, error) {

	orderID, err := os.qry.Checkout(UserID, newOrder, cartItems)

	if err != nil {
		log.Println("Checkout orders error", err.Error())
		return "", errors.New("interal server error")
	}
	totalPrice, err := os.qry.GetTotalOrderPrice(orderID)
	if err != nil {
		log.Println("Checkout payment error", err.Error())
		return "", errors.New("interal server error")
	}
	log.Print("================", totalPrice, "================")

	paymentURL, err := utils.Payment(orderID, uint64(totalPrice))
	if err != nil {
		log.Println("Checkout payment error", err.Error())
		return "", errors.New("interal server error")
	}

	currentOrder, err := os.qry.GetOrder(orderID)
	if err != nil {
		log.Println("Get Order qry error", err.Error())
		return "", errors.New("interal server error")
	}
	currentOrder.PaymentURL = paymentURL
	currentOrder.TotalOrder = uint64(totalPrice)

	err = os.qry.UpdateOrder(orderID, currentOrder)
	if err != nil {
		log.Println("Update Order qry error", err.Error())
		return "", errors.New("interal server error")
	}
	return paymentURL, nil
}
