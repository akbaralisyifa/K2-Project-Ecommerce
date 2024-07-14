package handler

import "ecommerce/internal/features/users"

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToModelUser(ur UserRequest) users.User{
	return users.User{
		Fullname: ur.Fullname,
		Email: ur.Email,
		Password: ur.Password,
	}
}