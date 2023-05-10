package main

import (
	"fmt"
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Connect()
	handler.InitRouting(e)
	// routes := e.Routes()
	// for _, route := range routes {
	// 	fmt.Printf("%s %s\n", route.Method, route.Path)
	// }

	for _, r := range e.Routes() {
		fmt.Printf("%s %s\n", r.Method, r.Path)
	}

	e.Logger.Fatal(e.Start(":8088"))
}
