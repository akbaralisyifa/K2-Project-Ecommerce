package service

import (
	"ecommerce/internal/features/users"
	"ecommerce/internal/utils"
	"errors"
	"log"
)

type UserServices struct {
	qry users.Query
	pwd utils.HashingPwInterface
	vld utils.ValidatorUtilityInterface
}

func NewUserService(q users.Query, p utils.HashingPwInterface, v utils.ValidatorUtilityInterface) users.Service {
	return &UserServices{
		qry: q,
		pwd: p,
		vld: v,
	}
};


func (us *UserServices) Register(newUser users.User) error {
	
	// validasi data
	err := us.vld.RegisterValidator(newUser.Username, newUser.Email, newUser.Password);
	if err != nil {
		log.Println("validator register error", err.Error());
		return err
	}

	// hashing password
	hashPw, err := us.pwd.GeneretePassword(newUser.Password);
	if err != nil {
		log.Println("register generete password error", err.Error());
		return	err
	}

	newUser.Password = string(hashPw);

	// register
	err = us.qry.Register(newUser);
	if err != nil {
		log.Println("register sql error:", err.Error());
		return errors.New("an error occurred on the server while processing data");
	}

	return nil;
}
