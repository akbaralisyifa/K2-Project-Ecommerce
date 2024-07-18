package repository

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"

	"gorm.io/gorm"
)

type OrderModels struct {
	db *gorm.DB
	crt cartitems.Query
};

func NewOrderModels(connect *gorm.DB, c cartitems.Query) orders.Query {
	return &OrderModels{
		db: connect,
		crt: c,
	}
};

// create orders
func (om *OrderModels) CreateOrders(newOrders orders.Order) (uint, error) {
	cnvData := ToOrderQuery(newOrders)
	err := om.db.Create(&cnvData).Error;

	if err != nil {
		return 0, err
	};

	return cnvData.ID, nil;
};

func (om *OrderModels) GetAllOrder(UserID uint) ([]orders.Order, error) {
	var result []Orders;

	err := om.db.Debug().Model(&Orders{}).Where("user_id = ? ", UserID).Preload("OrderItems").Find(&result).Error

	if err != nil {
		return []orders.Order{}, err;
	}

	return ToOrderEntityGetAll(result) , nil;
}

func (om *OrderModels) UpdateOrder(OrderID uint, updateOrder orders.Order) error {
	cnvData := ToOrderQuery(updateOrder)
	err := om.db.Model(&Orders{}).Where("order_id = ?", OrderID).Updates(&cnvData).Error

	if err != nil {
		return err
	};

	return nil
};

// Order Items
func(om *OrderModels) CreateOrderItems(orderID uint, newOrderItems []cartitems.CartItem) error {
	cnvData := ToOrderItemsQuery(orderID, newOrderItems)
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
// fungsi : - create order - create order item - delete carts - update product
func (om *OrderModels) Checkout(UserID uint, newOrder orders.Order, cartItems []cartitems.CartItem ) error {

	orderID, err := om.CreateOrders(newOrder);
	if err != nil {
		return err
	};

	err = om.CreateOrderItems(orderID, cartItems);
	if err != nil {
		return err
	};
	
	err = om.crt.DeleteCartItemByUserID(UserID)
	if err != nil {
		return err
	};

	// updateProduct products.Product
	// err = om.prd.UpdateProduct(productID, updateProduct);
	// if err != nil {
	// 	return err
	// };


	return nil;
}
