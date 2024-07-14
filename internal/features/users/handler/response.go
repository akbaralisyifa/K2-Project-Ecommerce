package handler

import "ecommerce/internal/features/users"

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToLoginResponse(input users.User) LoginResponse{
	return LoginResponse{
		ID: input.ID,
		Username: input.Username,
		Email: input.Email,
	}
}