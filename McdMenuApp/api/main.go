package main

import (
    "github.com/kunikida123456/McdMenuApp/config"
    "github.com/kunikida123456/McdMenuApp/infra"
    "github.com/kunikida123456/McdMenuApp/interface/handler"
    "github.com/kunikida123456/McdMenuApp/usecase"

    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/labstack/echo"
)

func main() {
    itemRepository := infra.NewItemRepository(config.NewDB())
    itemUsecase := usecase.NewItemUsecase(itemRepository)
    itemHandler := handler.NewItemHandler(itemUsecase)

    e := echo.New()
    handler.InitRouting(e, itemHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
