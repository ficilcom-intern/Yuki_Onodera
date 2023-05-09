package handler

import (
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/infra"
	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	meals := e.Group("/meals")
	mealHandler := NewMealHandler(usecase.NewMealUsecase(infra.NewMealRepository(config.DB)))

	meals.GET("/:id", mealHandler.Get)
	meals.POST("", mealHandler.Post)
	meals.PUT("/:id", mealHandler.Put)
	meals.DELETE("/:id", mealHandler.Delete)

	users := e.Group("/users")
	userHandler := NewUserHandler(usecase.NewUserUsecase(infra.NewUserRepository(config.DB)))

	users.POST("/signup", userHandler.Signup)
	users.POST("/login", userHandler.Login)
	users.POST("/logout", userHandler.Logout)

	return e

}
