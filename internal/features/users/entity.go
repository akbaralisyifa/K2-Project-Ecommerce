package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Query interface {
	Register(newUsers User) error
	Login(email string) (User, error)
}

type Service interface {
	Register(newUser User) error
	Login(email string, password string) (User,string, error)
}

type RegisterValidation struct {
	Username string `validate:"required,alpanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=7,alphanum"`
}

type LoginValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}