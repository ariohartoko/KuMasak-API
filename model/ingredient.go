package model

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primaryKey" form:"id"`
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
}
