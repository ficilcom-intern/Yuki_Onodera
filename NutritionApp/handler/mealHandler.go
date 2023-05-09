package handler

import (
	"net/http"
	"strconv"

	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo"
)

// MealHandler meal handlerのinterface
type MealHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type mealHandler struct {
	mealUsecase usecase.MealUsecase
}

// NewMealHandler meal handlerのコンストラクタ
func NewMealHandler(mealUsecase usecase.MealUsecase) MealHandler {
	return &mealHandler{mealUsecase: mealUsecase}
}

type requestMeal struct {
	Memo     string  `json:"memo" `
	Type string  `json:"type" `
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

type responseMeal struct {
	ID   int     `json:"id"`
	Type string  `json:"type"`
	Memo     string  `json:"memo"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

// Post mealを保存するときのハンドラー
func (mh *mealHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestMeal
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdMeal, err := mh.mealUsecase.Create(req.Memo, req.Type, req.Carbs, req.Fat, req.Protein, req.Calories)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMeal{
			ID:   createdMeal.ID,
			Type: createdMeal.Type,
			Memo:     createdMeal.Memo,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get mealを取得するときのハンドラー
func (th *mealHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundMeal, err := th.mealUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMeal{
			Memo:     foundMeal.Memo,
			Type:     foundMeal.Type,
			Carbs:    foundMeal.Carbs,
			Fat:      foundMeal.Fat,
			Protein:  foundMeal.Protein,
			Calories: foundMeal.Calories,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put mealを更新するときのハンドラー
func (th *mealHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestMeal
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedMeal, err := th.mealUsecase.Update(id, req.Memo, req.Type, req.Carbs, req.Fat, req.Protein, req.Calories)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMeal{
			ID:   updatedMeal.ID,
			Type: updatedMeal.Type,
			Memo:     updatedMeal.Memo,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete mealを削除するときのハンドラー
func (th *mealHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.mealUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
