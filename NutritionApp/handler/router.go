package handler

import (
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/infra"
	"kunikida123456/NutritionApp/usecase"
	"kunikida123456/NutritionApp/util"
	"kunikida123456/NutritionApp/myerror"

	echojwt "github.com/labstack/echo-jwt/v4"
	

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	meals := e.Group("/meals")

	mealHandler := NewMealHandler(usecase.NewMealUsecase(infra.NewMealRepository(config.DB())))
	meals.Use(echojwt.WithConfig(util.Config))
	// meals.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	meals.GET("/:id", mealHandler.Get)
	meals.POST("", mealHandler.Post)
	meals.PUT("/:id", mealHandler.Put)
	meals.DELETE("/:id", mealHandler.Delete)

	users := e.Group("/users")
	userHandler := NewUserHandler(usecase.NewUserUsecase(infra.NewUserRepository(config.DB())))

	users.POST("/signup", userHandler.Signup)
	users.POST("/login", userHandler.Login)
	users.POST("/logout", userHandler.Logout)

	e.Use(middleware.Logger())

	return e

}
