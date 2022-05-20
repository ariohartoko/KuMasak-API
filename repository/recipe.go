package repository

import (
	"fmt"
	"kumasak/domain"
	"kumasak/model"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//Count the bookmarks of a recipe
func (r *repoRecipe) bookmarkCount(id uint) int64 {
	var bookmark int64
	r.DB.Debug().
		Table("bookmarks").
		Where("recipe_id = ?", id).
		Count(&bookmark)
	return bookmark
}

//Quantity of an ingredient of a recipe
func (r *repoRecipe) ingredientCount(id, idi uint) float64 {
	var quantity float64
	r.DB.Debug().
		Table("recipe_ingredients").
		Select("quantity").
		Where("recipe_id = ? AND ingredient_id = ?", id, idi).
		Scan(&quantity)
	return quantity
}

//Get all recipes
func (r *repoRecipe) GetRecipes() []model.Recipe {
	var recipes []model.Recipe
	count := 0
	var wg sync.WaitGroup
	r.DB.Debug().
		Preload(clause.Associations).
		Find(&recipes)

	for _, v := range recipes {
		for range v.Ingredients {
			count++
		}
	}

	wg.Add(count + len(recipes))
	for i, v := range recipes {
		recipes[i].TotalPrice = 0
		for j, vi := range v.Ingredients {
			go func(i2, j2 int, v2 model.Recipe, vi2 model.Ingredient) {
				defer wg.Done()
				recipes[i2].Ingredients[j2].Quantity = r.ingredientCount(v2.ID, vi2.ID)
			}(i, j, v, vi)
		}
		go func(i3 int, v3 model.Recipe) {
			defer wg.Done()
			recipes[i3].Bookmarks = r.bookmarkCount(v3.ID)
		}(i, v)
	}
	wg.Wait()

	for i := range recipes {
		for _, v := range recipes[i].Ingredients {
			recipes[i].TotalPrice += v.Quantity * v.Price
		}
	}
	return recipes
}

//Get recipe by ID
func (r *repoRecipe) GetRecipeByID(id uint) (recipe model.Recipe, err error) {
	recipe = model.Recipe{}
	count := 0
	recipe.TotalPrice = 0
	var wg sync.WaitGroup

	if err = r.DB.Debug().
		Preload(clause.Associations).
		First(&recipe, id).
		Error; err != nil {
		return model.Recipe{}, err
	}
	for range recipe.Ingredients {
		count++
	}
	wg.Add(count)
	for i, v := range recipe.Ingredients {
		go func(i2 int, v2 model.Ingredient) {
			defer wg.Done()
			recipe.Ingredients[i2].Quantity = r.ingredientCount(recipe.ID, v2.ID)
		}(i, v)
	}
	wg.Wait()
	recipe.Bookmarks = r.bookmarkCount(recipe.ID)

	for _, v := range recipe.Ingredients {
		recipe.TotalPrice += v.Quantity * v.Price
	}

	return recipe, nil
}

//Get recipes by string
func (r *repoRecipe) GetRecipesByString(s string) []model.Recipe {
	recipes := []model.Recipe{}
	var wg sync.WaitGroup
	count := 0
	fmt.Println(s)
	words := strings.Fields(s)
	fmt.Println(words)
	r.DB.Debug().
		Where("title IN ?", words).
		Preload(clause.Associations).
		Find(&recipes)

	for _, v := range recipes {
		for range v.Ingredients {
			count++
		}
	}

	wg.Add(len(recipes))
	for i, v := range recipes {
		go func(i2 int, v2 model.Recipe) {
			defer wg.Done()
			recipes[i2], _ = r.GetRecipeByID(v2.ID)
		}(i, v)
	}
	wg.Wait()

	return recipes
}

//Create recipe
func (r *repoRecipe) CreateRecipe(recipe model.Recipe) (id uint, err error) {
	recipe.ViewCount = 0
	temp := recipe
	res := r.DB.Debug().Omit("ID", "TotalPrice", "Comments", "Bookmarks", "Ingredients", "Tags").Create(&recipe)
	if res.RowsAffected < 1 {
		return 0, fmt.Errorf("error creating user")
	}

	for _, v := range temp.Ingredients {
		r.DB.Debug().Exec("INSERT INTO `recipe_ingredients` (recipe_id, ingredient_id, quantity) VALUES ( ? , ? , ? )", recipe.ID, v.ID, v.Quantity)
	}

	for _, v := range temp.Tags {
		r.DB.Debug().Exec("INSERT INTO recipe_tags (recipe_id, tag_id) VALUES ( ? , ? )", recipe.ID, v.ID)
	}

	return recipe.ID, nil
}

//Update recipe
func (r *repoRecipe) UpdateRecipe(id uint, recipe model.Recipe) error {
	recipe.ID = id
	temp := recipe
	res := r.DB.Debug().Omit("TotalPrice", "Comments", "Bookmarks", "Ingredients", "Tags", "ViewCount").Save(&recipe)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error updating user")
	}

	for _, v := range temp.Ingredients {
		r.DB.Exec("INSERT INTO recipe_ingredients (recipe_id, ingredient_id) VALUES ( ? , ? )", recipe.ID, v.ID)
		r.DB.Exec("INSERT INTO recipe_ingredients (quantity) VALUES ( ? ) WHERE recipe_id = ? AND ingredient_id = ?", v.Quantity, recipe.ID, v.ID)
	}

	for _, v := range temp.Tags {
		r.DB.Exec("INSERT INTO recipe_tags (recipe_id, tag_id) VALUES ( ? , ? )", recipe.ID, v.ID)
	}

	return nil
}

//Delete recipe from database
func (r *repoRecipe) DeleteRecipe(id uint) error {
	recipe := model.Recipe{}
	recipe.ID = id
	res := r.DB.Find(&recipe)
	if res.RowsAffected < 1 {
		return fmt.Errorf("recipe not found")
	}

	r.DB.Exec("DELETE FROM bookmarks WHERE recipe_id = ?", id)
	r.DB.Exec("DELETE FROM comments WHERE recipe_id = ?", id)
	r.DB.Exec("DELETE FROM recipe_tags WHERE recipe_id = ?", id)
	r.DB.Exec("DELETE FROM recipe_ingredients WHERE recipe_id = ?", id)
	r.DB.Exec("DELETE FROM recipes WHERE id = ?", id)

	return nil
}

func NewRecipeRepo(db *gorm.DB) domain.RecipeRepoAdapter {
	return &repoRecipe{
		DB: db,
	}
}
