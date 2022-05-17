package main

import (
	"kumasak/config"
	"kumasak/route"

	"github.com/labstack/echo/v4"
	_ "gorm.io/gorm/clause"
)

func main() {
	conf := config.InitConfiguration()

	e := echo.New()

	route.HealthAPI(e, conf)
	route.RegisterRecipeGroupAPI(e, conf)

	e.Logger.Fatal(e.Start(config.InitConfiguration().SERVER_ADDRESS))

	/* recipe := model.Recipe{}
	DB.Debug().
		Preload(clause.Associations).
		Preload("Ingredients.Quantity").
		Find(&recipe)
	fmt.Print(recipe) */
}
