package repository

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/features/products"

	"gorm.io/gorm"
)

type OrderModels struct {
	db *gorm.DB
	crt cartitems.Query
	prd products.Query
};

func NewOrderModels(connect *gorm.DB, c cartitems.Query, p products.Query) orders.Query {
	return &OrderModels{
		db: connect,
		crt: c,
		prd: p,
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
func (om *OrderModels) Checkout(UserID uint, newOrder orders.Order, productID uint, updateProduct products.Product) error {
	
	cartItems, err := om.crt.GetAllCartItems(UserID);
	if err != nil {
		return err;
	}

	orderID, err := om.CreateOrders(newOrder);
	if err != nil {
		return err
	};

	err = om.CreateOrderItems(orderID, cartItems);
	if err != nil {
		return err
	};
	
	err = om.crt.DeleteCartItem(productID, UserID)
	if err != nil {
		return err
	};

	err = om.prd.UpdateProduct(productID, updateProduct);
	if err != nil {
		return err
	};


	return nil;
}
