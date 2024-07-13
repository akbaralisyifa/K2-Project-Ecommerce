package utils

import (
	"ecommerce/internal/features/users"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
);

type ValidatorUtilityInterface interface{
	RegisterValidator(username string, email string, password string) error
	LoginValidation(email string, password string) error
}

type ValidateUtility struct {
	vldt validator.Validate
};

func NewValidatorUtility(v validator.Validate) ValidatorUtilityInterface {
	return &ValidateUtility{
		vldt: v,
	}
}

func (vu *ValidateUtility) RegisterValidator(username string, email string, password string) error {
		err := vu.vldt.Struct(users.RegisterValidation{Username: username, Email: email, Password: password});

		if err != nil {
			log.Println("register validator error", err.Error());
			return errors.New("register validator error")
		};

		return nil;
};

func (vu *ValidateUtility) LoginValidation(email string, password string) error {
	err := vu.vldt.Struct(users.LoginValidation{Email: email, Password: password});

	if err != nil {
		log.Println("login validator error", err.Error());
		return errors.New("login validator error")
	};

	return	nil;
}