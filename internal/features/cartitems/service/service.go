package service

import (
	"ecommerce/internal/features/cartitems"
	"errors"
	"log"

	"gorm.io/gorm"
)

type CartItemServices struct {
	qry cartitems.Query

}

func NewCartItemService(q cartitems.Query, ) cartitems.Service {
	return &CartItemServices{
		qry: q,
	}
}

func (cs *CartItemServices) AddCartItem(newCartItem cartitems.CartItem, userID uint) error {

	// add Cart

	result, err := cs.qry.GetProduct(newCartItem.ProductID)
	if err != nil {
		log.Print("get product query error", err.Error())
		return errors.New("internal server error")
	}

	if result.Stock < int(newCartItem.Qty) {
		log.Print("Stock not enough for request")
		return errors.New("internal server error")
	}

	newCartItem.UserID = userID
	newCartItem.TotalPrice = uint(result.Price) * newCartItem.Qty

	_, err = cs.qry.GetCartItem(newCartItem.ProductID, newCartItem.UserID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			err = cs.qry.AddCartItem(newCartItem)
			if err != nil {
				log.Print("add cart item query error", err.Error())
				return errors.New("internal server error")
			}
			return nil
		}
		return errors.New("internal server error")
	}

	if int(newCartItem.Qty) < 1 {
		log.Print("item deleted")
		cs.qry.DeleteCartItem(newCartItem.ProductID, newCartItem.UserID)
		return nil
	}

	err = cs.qry.UpdateCartItem(newCartItem)
	if err != nil {
		log.Print("update Cart item query error", err.Error())
		return errors.New("internal server error")
	}
	return nil
}

func (cs *CartItemServices) DeleteCartItem(productID uint, userID uint) error {
	result, err := cs.qry.GetCartItem(productID, userID)
	if err != nil {
		log.Print("not found", err.Error())
		return errors.New("internal server error")
	}

	if result.UserID != userID {
		return errors.New(" unauthorize delete action")
	}
	err = cs.qry.DeleteCartItem(productID, userID)

	if err != nil {
		log.Print("delete Cart query error", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (cs *CartItemServices) GetAllCartItems(userID uint) ([]cartitems.CartItem, error) {

	result, err := cs.qry.GetAllCartItems(userID)
	msg := "internal server error"
	if err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "not found"
		}
		return []cartitems.CartItem{}, errors.New(msg)
	}

	return result, nil
}

func (cs *CartItemServices) DeleteCartItemByUserID(UserID uint) error {

	err := cs.qry.DeleteCartItemByUserID(UserID);
	msg := "internal server error"

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "not found"
		}
		return errors.New(msg)
	}

	return  nil
}
