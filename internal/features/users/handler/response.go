package handler

import "ecommerce/internal/features/users"

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID         uint   `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	ImgProfile string `json:"image_profile"`
}

func ToLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

func ToGetUserResponse(input users.User) UserResponse {
	return UserResponse{
		ID: 	  input.ID,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		ImgProfile: input.ImgProfile,
	}
}