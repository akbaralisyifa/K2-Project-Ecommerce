package repository

import (
	"ecommerce/internal/features/cartitems"

	"gorm.io/gorm"
)

type CartItemModels struct {
	db *gorm.DB
}

func NewCartItemModels(connect *gorm.DB) cartitems.Query {
	return &CartItemModels{
		db: connect,
	}
}

// Add CartItem
func (cm *CartItemModels) AddCartItem(newCartItems cartitems.CartItem) error {
	cnvData := ToCartItemsQuery(newCartItems)
	err := cm.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

// Get one data CartItem
func (cm *CartItemModels) GetCartItem(productID uint, userID uint) (cartitems.CartItem, error) {
	var result CartItems
	err := cm.db.Where(&CartItems{ProductID: productID, UserID: userID}).First(&result).Error

	if err != nil {
		return cartitems.CartItem{}, err
	}

	return result.ToCartItemsEntity(), nil
}

// update data CartItem
func (cm *CartItemModels) UpdateCartItem(updateCartItem cartitems.CartItem) error {
	cnvQuery := ToCartItemsQuery(updateCartItem)
	qry := cm.db.Where(&CartItems{ProductID: updateCartItem.ProductID, UserID: updateCartItem.UserID}).Updates(&cnvQuery)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cm *CartItemModels) DeleteCartItem(productID uint, userID uint) error {
	qry := cm.db.Where(&CartItems{ProductID: productID, UserID: userID}).Delete(&CartItems{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cm *CartItemModels) GetAllCartItems(userID uint) ([]cartitems.CartItem, error) {
	var result []CartItems
	var resultConvert []cartitems.CartItem
	err := cm.db.Where("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return []cartitems.CartItem{}, err
	}
	for _, v := range result {
		resultConvert = append(resultConvert, v.ToCartItemsEntity())
	}
	return resultConvert, nil
}

// Get one data product
func (cm *CartItemModels) GetProduct(productID uint) (cartitems.Product, error) {
	var result Products
	err := cm.db.Where("id = ?", productID).First(&result).Error

	if err != nil {
		return cartitems.Product{}, err
	}

	return result.ToProductsEntity(), nil
};

// delete cart item by userid 
func (cm *CartItemModels) DeleteCartItemByUserID(UserID uint) error {

	qry := cm.db.Where("user_id = ?", UserID).Delete(&CartItems{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
}
