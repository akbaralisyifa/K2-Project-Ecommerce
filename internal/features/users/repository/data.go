package repository

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username 	string
	Email 		string
	Password 	string
}