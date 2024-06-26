package db

import "gorm.io/gorm"

type ClientUsers struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
