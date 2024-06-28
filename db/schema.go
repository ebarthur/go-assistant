package db

import "gorm.io/gorm"

type ClientUsers struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	Histories []History `gorm:"foreignKey:UserID"`
}

type History struct {
	gorm.Model
	UserID   uint
	Request  string
	Response string
	Endpoint string
}
