package service

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/utils"
	"errors"
	"log"

	"gorm.io/gorm"
)

type CartItemServices struct {
	qry cartitems.Query
	pwd utils.HashingPwInterface
	vld utils.ValidatorUtilityInterface
	jwt utils.JwtUtilityInterface
}

func NewCartItemService(q cartitems.Query, p utils.HashingPwInterface, v utils.ValidatorUtilityInterface, j utils.JwtUtilityInterface) cartitems.Service {
	return &CartItemServices{
		qry: q,
		pwd: p,
		vld: v,
		jwt: j,
	}
}

func (cs *CartItemServices) AddCartItem(newCartItem cartitems.CartItem, userID uint) error {

	// add Cart
	newCartItem.UserID = userID
	_, err := cs.qry.GetCartItem(newCartItem.ProductID, newCartItem.UserID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			err = cs.qry.AddCartItem(newCartItem)
			if err != nil {
				log.Print("add cart item query error", err.Error())
				return errors.New("internal server error")
			}
		}
		return errors.New("internal server error")
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
