package repository

import (
	"ecommerce/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Fullname 	string
	Email 		string
	Password 	string
	Phone		string
	Address		string
};

// dari database di pindah ke entity
func (u *Users) ToUsersEntity() users.User{
	return users.User{
		ID: 	  u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Phone: 	  u.Phone,
		Address:  u.Address,
	}
}

// dari entity pindah ke database
func ToUsersQuery(input users.User) Users {
	return Users{
		Fullname: input.Fullname,
		Email: 	  input.Email,
		Password: input.Password,
		Phone:	  input.Phone,
		Address:  input.Address,
	}
}
