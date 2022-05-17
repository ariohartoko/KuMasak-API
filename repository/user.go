package repository

/*
import (
	"kumasak/model"
	"sync"

	"gorm.io/gorm/clause"
	//"gorm.io/gorm"
	//"gorm.io/gorm/clause"
)
*/
/*
//Insert a new user
func (r *repoRecipe) CreateUser(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error creating user")
	}
	return nil
}
*/

/*
//Get all users
func (r *repoUser) GetUsers() []model.User {
	users := []model.User{}
	r.DB.Debug().
		Preload(clause.Associations).
		Preload("recipes.ingredients").
		Preload("recipes.tags").
		Find(&users)

	return users

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
*/

/*
//Get top user by ID
func (r *repoMysql) GetUserByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}
	return user, err
}

//Get top user by Name
func (r *repoMysql) GetUserByName(name string) (user model.User, err error) {
	res := r.DB.Where("name = ?", name).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}
	return user, err
}

//Get top user by Email
func (r *repoMysql) GetUserByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}
	return user, err
}

//Update user by ID
func (r *repoMysql) UpdateUserByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error updating user")
	}

	return nil
}

//Delete user by ID
func (r *repoMysql) DeleteByID(id int) error {
	res := r.DB.Delete(&model.User{
		//ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}
*/
