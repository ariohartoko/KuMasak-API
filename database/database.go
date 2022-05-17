package database

import (
	"fmt"
	"kumasak/config"
	"kumasak/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		fmt.Println("error opening connection : ", err)
	}

	err = DB.SetupJoinTable(&model.Recipe{}, "Ingredients", &model.RecipeIngredient{})
	if err != nil {
		fmt.Print("error setting up join table : ", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Recipe{}, &model.Tag{}, &model.Ingredient{}, &model.Comment{})
	if err != nil {
		fmt.Print("error migrating table : ", err)
	}

	return DB
}
