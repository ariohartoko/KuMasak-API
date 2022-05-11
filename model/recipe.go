package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey" form:"id"`
	User_ID     int    `json:"user_id" form:"user_id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Steps       string `json:"steps" form:"steps"`
	View_count  int    `json:"view_count" form:"view_count"`
}
