package service

import (
	"kumasak/config"
	"kumasak/domain"
	"kumasak/model"
)

type recipeService struct {
	conf config.Config
	repo domain.RecipeRepoAdapter
}

func (rs *recipeService) GetRecipesService() []model.Recipe {
	return rs.repo.GetRecipes()
}

func (rs *recipeService) GetRecipeByIDService(id uint) (recipe model.Recipe, err error) {
	return rs.repo.GetRecipeByID(id)
}

func (rs *recipeService) GetRecipesByStringService(s string) []model.Recipe {
	return rs.repo.GetRecipesByString(s)
}

func (rs *recipeService) CreateRecipeService(recipe model.Recipe) (id uint, err error) {
	return rs.repo.CreateRecipe(recipe)
}

func (rs *recipeService) UpdateRecipeService(id uint, recipe model.Recipe) error {
	return rs.repo.UpdateRecipe(id, recipe)
}

func (rs *recipeService) DeleteRecipeService(id uint) error {
	return rs.repo.DeleteRecipe(id)
}

func NewServiceRecipe(repo domain.RecipeRepoAdapter, conf config.Config) domain.ServiceAdapter {
	return &recipeService{
		repo: repo,
		conf: conf,
	}
}
