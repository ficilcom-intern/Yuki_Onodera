package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, itemHandler ItemHandler) {
	e.GET("/item", itemHandler.Get())
}
