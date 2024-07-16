package users

import "github.com/labstack/echo/v4"

type User struct {
	ID         uint
	Fullname   string
	Email      string
	Password   string
	Phone	   string
	Address	   string
	ImgProfile string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type Query interface {
	Register(newUsers User) error
	Login(email string) (User, error)
	GetUser(ID uint)(User, error)
	UpdateUser(ID uint, updateUser User) error
	DeleteUser(ID uint) error
}

type Service interface {
	Register(newUser User) error
	Login(email string, password string) (User,string, error)
	GetUser(ID uint) (User, error)
	UpdateUser(ID uint, updateUser User) error
	DeleteUser(ID uint) error
}

type RegisterValidation struct {
	Fullname string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}