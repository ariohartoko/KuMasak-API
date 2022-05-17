package route

import (
	"kumasak/config"
	"kumasak/model"

	"github.com/labstack/echo/v4"
)

func HealthAPI(e *echo.Echo, conf config.Config) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, model.Result{
			Code:    200,
			Message: "sehat gan",
		})
	})
}
