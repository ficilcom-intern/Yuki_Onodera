package main

import (
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Connect()
	handler.InitRouting(e)
	e.Logger.Fatal(e.Start(":8088"))
}
