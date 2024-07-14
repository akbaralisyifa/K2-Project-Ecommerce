package repository

import (
	"ecommerce/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username 	string
	Email 		string
	Password 	string
};

// dari database di pindah ke entity
func (u *Users) ToUsersEntity() users.User{
	return users.User{
		ID: 	  u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

// dari entity pindah ke database
func ToUsersQuery(input users.User) Users {
	return Users{
		Username: input.Username,
		Email: 	  input.Email,
		Password: input.Password,
	}
}