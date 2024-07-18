package service

import (
	"ecommerce/internal/features/users"
	"ecommerce/internal/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	qry users.Query
	pwd utils.HashingPwInterface
	vld utils.ValidatorUtilityInterface
	jwt utils.JwtUtilityInterface
}

func NewUserService(q users.Query, p utils.HashingPwInterface, v utils.ValidatorUtilityInterface, j utils.JwtUtilityInterface) users.Service {
	return &UserServices{
		qry: q,
		pwd: p,
		vld: v,
		jwt: j,
	}
}

func (us *UserServices) Register(newUser users.User) error {
	// validasi data
	err := us.vld.RegisterValidator(newUser.Fullname, newUser.Email, newUser.Password)
	if err != nil {
		log.Println("validator register error", err.Error())
		return err
	}

	// hashing password
	hashPw, err := us.pwd.GeneretePassword(newUser.Password)
	if err != nil {
		log.Println("register generete password error", err.Error())
		return err
	}

	newUser.Password = string(hashPw)

	// register
	err = us.qry.Register(newUser)
	if err != nil {
		log.Println("register sql error:", err.Error())
		return errors.New("error in server")
	}

	return nil
}

func (us *UserServices) Login(email string, password string) (users.User, string, error) {

	err := us.vld.LoginValidation(email, password)
	// Jika validasi gagal
	if err != nil {
		log.Println("validation error:", err.Error())
		return users.User{}, "", err
	}

	result, err := us.qry.Login(email)
	if err != nil {
		log.Println("login sql error: ", err.Error())
		return users.User{}, "", errors.New("error in server")
	}

	// cek password
	err = us.pwd.CheckPassword([]byte(password), []byte(result.Password))
	if err != nil {
		log.Println("Error On Password", err)
		return users.User{}, "", errors.New(bcrypt.ErrMismatchedHashAndPassword.Error())
	}

	// generete token
	token, err := us.jwt.GenereteJwt(result.ID)
	if err != nil {
		log.Println("Error On Jwt ", err)
		return users.User{}, "", errors.New("error on JWT")
	}

	return result, token, nil
}

func (us *UserServices) GetUser(ID uint) (users.User, error) {
	result, err := us.qry.GetUser(ID)

	if err != nil {
		log.Print("get user query error", err.Error())
		return users.User{}, errors.New("internal server error")
	}

	return result, nil
}

func (us *UserServices) UpdateUser(ID uint, updateUser users.User) error {

	if updateUser.Password != "" {
		hashPw, err := us.pwd.GeneretePassword(updateUser.Password)

		if err != nil {
			log.Print("update generete password error", err.Error())
			return errors.New("update generete password error")
		}

		updateUser.Password = string(hashPw)
	}
	if updateUser.Password != "" {
		hashPw, err := us.pwd.GeneretePassword(updateUser.Password)
		if err != nil {
			log.Print("update generete password error", err.Error())
			return errors.New("update generete password error")
		}
		updateUser.Password = string(hashPw)
	}

	// update user
	err := us.qry.UpdateUser(ID, updateUser)

	if err != nil {
		log.Print("update user query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}

func (us *UserServices) DeleteUser(ID uint) error {
	err := us.qry.DeleteUser(ID)

	if err != nil {
		log.Print("delte user query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}
