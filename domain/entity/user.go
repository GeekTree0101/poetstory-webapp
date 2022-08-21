package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string
	Bio      string
}
