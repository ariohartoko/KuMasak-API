package route

import (
	"kumasak/config"
	"kumasak/controller"
	"kumasak/database"
	"kumasak/repository"
	"kumasak/service"

	"github.com/labstack/echo/v4"
)

func RegisterRecipeGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewRecipeRepo(db)

	svc := service.NewServiceRecipe(repo, conf)

	cont := controller.ServiceController{
		Sa: svc,
	}

	apiRecipe := e.Group("/recipe")

	apiRecipe.GET("/all", cont.GetRecipesController)
	apiRecipe.GET("/:id", cont.GetRecipeByIDController)
	apiRecipe.GET("/search/:search", cont.GetRecipesByStringController)
	apiRecipe.POST("", cont.CreateRecipeController)
	apiRecipe.PUT("/:id", cont.UpdateRecipeController)
	apiRecipe.DELETE("/:id", cont.DeleteRecipeController)
}
