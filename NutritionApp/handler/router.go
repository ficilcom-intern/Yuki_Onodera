package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, mealHandler MealHandler) {

	e.POST("/meal", mealHandler.Post())
	e.GET("/meal/:id", mealHandler.Get())
	e.PUT("/meal/:id", mealHandler.Put())
	e.DELETE("/meal/:id", mealHandler.Delete())

}
