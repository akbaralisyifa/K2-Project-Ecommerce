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

type GetUpdateRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`	
	Phone 	 string `json:"phone"`	
	Address  string `json:"address"`	
}

func ToModelUser(ur UserRequest) users.User{
	return users.User{
		Fullname: ur.Fullname,
		Email: 	  ur.Email,
		Password: ur.Password,
	}
};

func ToRequertUser(ur GetUpdateRequest) users.User {
	return users.User{
		Fullname: ur.Fullname,
		Email:    ur.Email,
		Password: ur.Password,
		Phone:    ur.Phone,
		Address:  ur.Address,
	}
}