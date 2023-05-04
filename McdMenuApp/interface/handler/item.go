package handler

import (
	"net/http"
	"strconv"

	"github.com/kunikida123456/McdMenuApp/usecase"

	"github.com/labstack/echo"
)

// ItemHandler item handlerのinterface
type ItemHandler interface {
	Get() echo.HandlerFunc
}

type itemHandler struct {
	itemUsecase usecase.ItemUsecase
}

// NewItemHandler item handlerのコンストラクタ
func NewItemHandler(itemUsecase usecase.ItemUsecase) ItemHandler {
	return &itemHandler{itemUsecase: itemUsecase}
}

type responseItem struct {
	Drink   string `json:"drink"`
	Burger  string `json:"burger"`
	Side    string `json:"side"`
	Barista string `json:"barista"`
}

// Get itemを取得するときのハンドラー
func (ih *itemHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundItem, err := ih.itemUsecase.MakeMenu()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			Burger:  foundItem.Burger,
			Drink:   foundItem.Drink,
			Side:    foundItem.Side,
			Barista: foundItem.Barista,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put itemを更新するときのハンドラー
func (ih *itemHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestItem
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedItem, err := ih.itemUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			ID:      updatedItem.ID,
			Title:   updatedItem.Title,
			Content: updatedItem.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete itemを削除するときのハンドラー
func (ih *itemHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = ih.itemUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
