package domain

import "kumasak/model"

type ServiceAdapter interface {
	GetRecipesService() []model.Recipe
	GetRecipeByIDService(id uint) (recipe model.Recipe, err error)
	GetRecipesByStringService(s string) []model.Recipe
	CreateRecipeService(recipe model.Recipe) (id uint, err error)
	UpdateRecipeService(id uint, recipe model.Recipe) error
	DeleteRecipeService(id uint) error
}
