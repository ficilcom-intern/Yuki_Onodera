package main

import (
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/handler"
	"kunikida123456/NutritionApp/infra"
	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo"
)

func main() {
	mealRepository := infra.NewMealRepository(config.NewDB())
	mealUsecase := usecase.NewMealUsecase(mealRepository)
	mealHandler := handler.NewMealHandler(mealUsecase)

	e := echo.New()
	handler.InitRouting(e, mealHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
