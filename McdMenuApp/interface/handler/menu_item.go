package handler

import (
	"net/http"
	"strconv"

	"github.com/kunikida123456/McdMenuApp/usecase"

	"github.com/labstack/echo"
)

// ItemHandler item handlerのinterface
type ItemHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type itemHandler struct {
	itemUsecase usecase.ItemUsecase
}

// NewItemHandler item handlerのコンストラクタ
func NewItemHandler(itemUsecase usecase.ItemUsecase) ItemHandler {
	return &itemHandler{itemUsecase: itemUsecase}
}

type responseItem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Get itemを取得するときのハンドラー
func (th *itemHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundItem, err := th.itemUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			ID:      foundItem.ID,
			Title:   foundItem.Title,
			Content: foundItem.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}
