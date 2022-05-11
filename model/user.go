package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Photo    string `json:"photo" form:"photo"`
}
