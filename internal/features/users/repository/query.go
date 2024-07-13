package repository

import (
	"ecommerce/internal/features/users"

	"gorm.io/gorm"
)

type UserModels struct {
	db *gorm.DB
};

func NewUserModels(connect *gorm.DB) users.Query {
	return &UserModels{
		db: connect,
	}
}

// Register
func (um *UserModels) Register(newUsers users.User)(error) {
	cnvData := ToUsersQuery(newUsers);
	err := um.db.Create(cnvData).Error

	if err != nil {
		return err
	}

	return nil;
}

// Login
func (um *UserModels) Login(email string) (users.User, error) {
	var result Users;
	err := um.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUsersEntity(), nil;
}