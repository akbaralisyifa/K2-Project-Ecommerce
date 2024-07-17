package repository

import (
	"ecommerce/internal/features/orders"

	"gorm.io/gorm"
)

type OrderModels struct {
	db *gorm.DB
};

func NewOrderModels(connect *gorm.DB) orders.Query {
	return &OrderModels{
		db: connect,
	}
};

// create orders
func (om *OrderModels) CreateOrders(newOrders orders.Order) error {
	cnvData := ToOrderQuery(newOrders)
	err := om.db.Create(&cnvData).Error;

	if err != nil {
		return err
	};

	return nil;
};

func (om *OrderModels) GetAllOrder(UserID uint) ([]orders.Order, error) {
	var result []Orders;

	err := om.db.Debug().Model(&Orders{}).Preload("OrderItems").Where("user_id = ? ", UserID).Find(&result).Error

	if err != nil {
		return []orders.Order{}, err;
	}

	return ToOrderEntityGetAll(result) , nil;
}

func (om *OrderModels) UpdateOrderStatus(OrderID uint, newStatus string) error {
	
	err := om.db.Model(&Orders{}).Where("order_id = ?", OrderID).Update("status", newStatus).Error;

	if err != nil {
		return err
	};

	return nil
};

// Order Items
func(om *OrderModels) CreateOrderItems(newOrderItems orders.OrderItems) error {
	cnvData := ToOrderItemsQuery(newOrderItems);
	err := om.db.Create(&cnvData).Error

	if err != nil {
		return err
	};

	return nil;
};

func (om *OrderModels) GetOrderItems(OrderID uint) ([]orders.OrderItems, error){
	var result []OrderItems
	err := om.db.Where("order_id = ?", OrderID).Find(&result).Error

	if err != nil {
		return []orders.OrderItems{}, err;
	}

	return ToOrderItemsGetAll(result), nil;
};


// fungsi checkout
// fungsi : - create order - create order item - delete carts
func (om *OrderModels) Checkout(newOrder orders.Order, newOrderItem orders.OrderItems) error {
	err := om.CreateOrders(newOrder);

	if err != nil {
		return err
	}

	err = om.CreateOrderItems(newOrderItem);

	if err != nil {
		return err
	};

	return nil;
}
