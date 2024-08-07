package repository

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/features/products"

	"gorm.io/gorm"
)

type OrderModels struct {
	db  *gorm.DB
	crt cartitems.Query
	prd products.Query
}

func NewOrderModels(connect *gorm.DB, c cartitems.Query, p products.Query) orders.Query {
	return &OrderModels{
		db:  connect,
		crt: c,
		prd: p,
	}
}

// create orders
func (om *OrderModels) CreateOrders(newOrders orders.Order) (uint, error) {
	cnvData := ToOrderQuery(newOrders)
	err := om.db.Create(&cnvData).Error

	if err != nil {
		return 0, err
	}

	return cnvData.ID, nil
}

func (om *OrderModels) GetAllOrder(UserID uint) ([]orders.Order, error) {
	var result []Orders

	err := om.db.Debug().Model(&Orders{}).Where("user_id = ? ", UserID).Preload("OrderItems").Find(&result).Error

	if err != nil {
		return []orders.Order{}, err
	}

	return ToOrderEntityGetAll(result), nil
}

func (om *OrderModels) GetAllOrderHistory(userID uint, orderID uint) ([]orders.Order, error) {
	var result []Orders

	err := om.db.Debug().Model(&Orders{}).Where("user_id = ? AND id = ?", userID, orderID).Preload("OrderItems").Find(&result).Error
	// err := om.db.Raw(qry, userID).Scan(&result).Error

	if err != nil {
		return []orders.Order{}, err
	}

	return ToOrderEntityGetAll(result), nil
}

func (om *OrderModels) UpdateOrder(OrderID uint, updateOrder orders.Order) error {
	cnvData := ToOrderQuery(updateOrder)
	err := om.db.Model(&Orders{}).Where("id = ?", OrderID).Updates(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

// Order Items
func (om *OrderModels) CreateOrderItems(orderID uint, newOrderItems []cartitems.CartItem) error {
	cnvData := ToOrderItemsQuery(orderID, newOrderItems)
	err := om.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

func (om *OrderModels) GetOrderItems(OrderID uint) ([]orders.OrderItems, error) {
	var result []OrderItems
	err := om.db.Where("order_id = ?", OrderID).Find(&result).Error

	if err != nil {
		return []orders.OrderItems{}, err
	}

	return ToOrderItemsGetAll(result), nil
}

// fungsi checkout
// fungsi : - create order - create order item - delete carts - update product
func (om *OrderModels) Checkout(UserID uint, newOrder orders.Order, cartItems []cartitems.CartItem) (uint, error) {
	tx := om.db.Begin()

	orderID, err := om.CreateOrders(newOrder)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = om.CreateOrderItems(orderID, cartItems)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = om.crt.DeleteCartItemByUserID(UserID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// updateProduct
	for _, val := range cartItems {
		product, err := om.prd.GetProduct(val.ProductID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		product.Stock -= int(val.Qty)
		err = om.prd.UpdateProduct(product.ID, product)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return orderID, nil
}

func (om *OrderModels) GetOrder(orderID uint) (orders.Order, error) {
	var result Orders

	err := om.db.Where("id = ?", orderID).Find(&result).Error

	if err != nil {
		return orders.Order{}, err
	}

	return result.ToOrderEntity(), nil
}

func (om *OrderModels) GetTotalOrderPrice(orderID uint) (int, error) {
	var result []OrderItems
	var totalPrice int
	err := om.db.Where("order_id = ?", orderID).Find(&result).Error

	if err != nil {
		return totalPrice, err
	}
	for _, v := range result {
		totalPrice = totalPrice + v.TotalPrice
	}
	return totalPrice, nil
}
