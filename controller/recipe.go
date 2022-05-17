package controller

import (
	"kumasak/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//Get all recipes
func (sc *ServiceController) GetRecipesController(c echo.Context) error {
	recipes := sc.Sa.GetRecipesService()

	return c.JSONPretty(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "Success getting all recipes",
		Data:    recipes,
	}, "	")
}

//Get a recipe by ID
func (sc *ServiceController) GetRecipeByIDController(c echo.Context) error {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	res, err := sc.Sa.GetRecipeByIDService(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Result{
			Code:    http.StatusNotFound,
			Message: "recipe not found",
			Data:    nil,
		})
	}

	return c.JSONPretty(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "success getting recipe",
		Data:    res,
	}, "	")
}

//Get recipes by string
func (sc *ServiceController) GetRecipesByStringController(c echo.Context) error {
	s := c.Param("search")
	res := sc.Sa.GetRecipesByStringService(s)
	if res == nil {
		return c.JSON(http.StatusNotFound, model.Result{
			Code:    http.StatusNotFound,
			Message: "no recipe found",
			Data:    nil,
		})
	}

	return c.JSONPretty(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "success getting recipes",
		Data:    res,
	}, "	")
}

//Create recipe
func (sc *ServiceController) CreateRecipeController(c echo.Context) error {
	recipe := model.Recipe{}
	c.Bind(&recipe)

	id, err := sc.Sa.CreateRecipeService(recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{
			Code:    http.StatusInternalServerError,
			Message: "failed creating recipe",
			Data:    nil,
		})
	}

	recipe = model.Recipe{}
	recipe, err = sc.Sa.GetRecipeByIDService(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{
			Code:    http.StatusInternalServerError,
			Message: "success creating recipe but failed getting recipe",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "Success creating user",
		Data:    recipe,
	})

}

//Update recipe by ID
func (sc *ServiceController) UpdateRecipeController(c echo.Context) error {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = sc.Sa.GetRecipeByIDService(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Result{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	recipe := model.Recipe{}
	c.Bind(&recipe)

	err = sc.Sa.UpdateRecipeService(uint(uintID), recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{
			Code:    http.StatusInternalServerError,
			Message: "failed updating recipe",
			Data:    nil,
		})
	}

	recipe = model.Recipe{}
	recipe, err = sc.Sa.GetRecipeByIDService(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{
			Code:    http.StatusInternalServerError,
			Message: "success updating recipe but failed getting recipe",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "Success updating recipe",
		Data:    recipe,
	})

}

//Delete recipe by ID
func (sc *ServiceController) DeleteRecipeController(c echo.Context) error {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = sc.Sa.GetRecipeByIDService(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Result{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = sc.Sa.DeleteRecipeService(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSONPretty(http.StatusOK, model.Result{
		Code:    http.StatusOK,
		Message: "success deleting recipes",
		Data:    nil,
	}, "	")
}
