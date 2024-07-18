package handler

import "ecommerce/internal/features/users"

type UserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUpdateRequest struct {
	Fullname   string `json:"fullname" form:"fullname"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Phone      string `json:"phone" form:"phone"`
	Address    string `json:"address" form:"address"`
	ImgProfile string
}

func ToModelUser(ur UserRequest) users.User {
	return users.User{
		Fullname: ur.Fullname,
		Email:    ur.Email,
		Password: ur.Password,
	}
}

func ToRequertUser(ur GetUpdateRequest) users.User {
	return users.User{
		Fullname:   ur.Fullname,
		Email:      ur.Email,
		Password:   ur.Password,
		Phone:      ur.Phone,
		Address:    ur.Address,
		ImgProfile: ur.ImgProfile,
	}
}
