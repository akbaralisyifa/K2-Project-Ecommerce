package utils

import (
	"ecommerce/internal/features/products"
	"ecommerce/internal/features/users"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type ValidatorUtilityInterface interface {
	RegisterValidator(username string, email string, password string) error
	LoginValidation(email string, password string) error
	AddProductValidation(name string, price int, Stock int) error
}

type ValidateUtility struct {
	vldt validator.Validate
}

func NewValidatorUtility(v validator.Validate) ValidatorUtilityInterface {
	return &ValidateUtility{
		vldt: v,
	}
}

func (vu *ValidateUtility) RegisterValidator(fullname string, email string, password string) error {
	err := vu.vldt.Struct(users.RegisterValidation{Fullname: fullname, Email: email, Password: password})

	if err != nil {
		log.Println("register validator error", err.Error())
		return errors.New("register validator error")
	}

	return nil
}

func (vu *ValidateUtility) LoginValidation(email string, password string) error {
	err := vu.vldt.Struct(users.LoginValidation{Email: email, Password: password})

	if err != nil {
		log.Println("login validator error", err.Error())
		return errors.New("login validator error")
	}

	return nil
}

func (vu *ValidateUtility) AddProductValidation(name string, price int, stock int) error {
	err := vu.vldt.Struct(products.AddProductValidation{Name: name, Price: price, Stock: stock})

	if err != nil {
		log.Println("add product validator error", err.Error())
		return errors.New("add product validator error")
	}

	return nil
}
