package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password,omitempty" form:"password"`
	Photo     string    `json:"photo" form:"photo"`
	Recipes   []Recipe  `json:"recipes"`
	Bookmarks []Recipe  `gorm:"many2many:bookmarks;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments  []Comment `json:"comments"`
}

type Recipe struct {
	gorm.Model
	UserID      uint         `json:"user_id" form:"user_id" gorm:"constraint:OnUpdateCASCADE,OnDelete:CASCADE"`
	Title       string       `json:"title" form:"title"`
	Description string       `json:"description" form:"description"`
	Ingredients []Ingredient `json:"ingredient" form:"ingredient" gorm:"many2many:recipe_ingredient;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TotalPrice  float64      `json:"total_price" gorm:"-"`
	Steps       string       `json:"steps" form:"steps"`
	ViewCount   int          `json:"view_count" gorm:"default:0"`
	Tags        []Tag        `json:"tag" form:"tag" gorm:"many2many:recipe_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments    []Comment    `json:"comment" form:"comment"`
	Bookmarks   int64        `gorm:"-" json:"bookmarks"`
}

type Comment struct {
	gorm.Model
	RecipeID uint   `json:"recipe_id" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID   uint   `json:"user_id" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comment  string `json:"comment" form:"comment"`
}

type Ingredient struct {
	ID          uint    `json:"id" gorm:"primaryKey" form:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price" gorm:"default:0"`
	Measurement string  `json:"measurement" form:"measurement"`
	Quantity    float64 `json:"quantity" form:"quantity" gorm:"-"`
}

type Tag struct {
	ID      uint     `gorm:"primaryKey" json:"id" form:"id"`
	Name    string   `json:"name" form:"name"`
	Recipes []Recipe `json:"recipe,omitempty" gorm:"many2many:recipe_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type RecipeIngredient struct {
	RecipeID     uint    `json:"-" gorm:"primaryKey"`
	IngredientID uint    `json:"-" gorm:"primaryKey"`
	Quantity     float64 `json:"quantity" form:"quantity"`
}
