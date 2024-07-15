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
	err := um.db.Create(&cnvData).Error

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
};

// Get one data user
func (um *UserModels) GetUser(ID uint)(users.User, error) {
	var result Users;
	err := um.db.Where("id = ?", ID).First(&result).Error;

	if err != nil {
		return users.User{}, err;
	};

	return result.ToUsersEntity(), nil;
};

// update data user
func (um *UserModels) UpdateUser(ID uint, updateUser users.User) error {
	cnvQuery := ToUsersQuery(updateUser);


	if updateUser.Fullname != "" {
		cnvQuery.Fullname = updateUser.Fullname
	}

	if updateUser.Email != "" {
		cnvQuery.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		cnvQuery.Password = updateUser.Password
	}

	if updateUser.Phone != ""{
		cnvQuery.Phone = updateUser.Phone
	}

	if updateUser.Address != "" {
		cnvQuery.Address = updateUser.Address
	}

	qry := um.db.Where("id = ?", ID).Updates(&cnvQuery);

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
};

func (um *UserModels) DeleteUser(ID uint) error {
	qry := um.db.Where("id = ?", ID).Delete(&Users{});

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
}