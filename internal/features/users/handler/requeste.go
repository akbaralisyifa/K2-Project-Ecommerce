package handler

import "ecommerce/internal/features/users"

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToModelUser(ur UserRequest) users.User{
	return users.User{
		Username: ur.Username,
		Email: ur.Email,
		Password: ur.Password,
	}
}