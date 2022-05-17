package domain

import (
	"kumasak/model"
)

type RecipeRepoAdapter interface {
	GetRecipes() []model.Recipe
	GetRecipeByID(id uint) (recipe model.Recipe, err error)
	GetRecipesByString(s string) []model.Recipe
	CreateRecipe(recipe model.Recipe) (id uint, err error)
	UpdateRecipe(id uint, recipe model.Recipe) error
	DeleteRecipe(id uint) error
}
