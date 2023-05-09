package handler

import (
	"net/http"
	"strconv"

	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo"
)

// MealHandler meal handlerのinterface
type MealHandler interface {
	Post(c echo.Context) error
	Get(c echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
}

type mealHandler struct {
	mealUsecase usecase.MealUsecase
}

// NewMealHandler meal handlerのコンストラクタ
func NewMealHandler(mealUsecase usecase.MealUsecase) MealHandler {
	return &mealHandler{mealUsecase: mealUsecase}
}

// Post mealを保存するときのハンドラー
func (mh *mealHandler) Post(c echo.Context) error {

	type (
		request struct {
			Memo     string  `json:"memo" `
			Type     string  `json:"type" `
			Carbs    float64 `json:"carbs"`
			Fat      float64 `json:"fat"`
			Protein  float64 `json:"protein"`
			Calories float64 `json:"calories"`
		}

		response struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
			Memo string `json:"memo"`
		}
	)

	req := new(request)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdMeal, err := mh.mealUsecase.Create(req.Memo, req.Type, req.Carbs, req.Fat, req.Protein, req.Calories)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := response{
		ID:   createdMeal.ID,
		Type: createdMeal.Type,
		Memo: createdMeal.Memo,
	}

	return c.JSON(http.StatusCreated, res)
}

// Get mealを取得するときのハンドラー
func (mh *mealHandler) Get(c echo.Context) error {

	type response struct {
		ID       int     `json:"id"`
		Memo     string  `json:"memo"`
		Type     string  `json:"type"`
		Carbs    float64 `json:"carbs"`
		Fat      float64 `json:"fat"`
		Protein  float64 `json:"protein"`
		Calories float64 `json:"calories"`
	}

	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundMeal, err := mh.mealUsecase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := response{
		ID:       foundMeal.ID,
		Memo:     foundMeal.Memo,
		Type:     foundMeal.Type,
		Carbs:    foundMeal.Carbs,
		Fat:      foundMeal.Fat,
		Protein:  foundMeal.Protein,
		Calories: foundMeal.Calories,
	}
	return c.JSON(http.StatusOK, res)

}

// Put mealを更新するときのハンドラー
func (mh *mealHandler) Put(c echo.Context) error {
	type (
		request struct {
			Memo     string  `json:"memo" `
			Type     string  `json:"type" `
			Carbs    float64 `json:"carbs"`
			Fat      float64 `json:"fat"`
			Protein  float64 `json:"protein"`
			Calories float64 `json:"calories"`
		}
		response struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
			Memo string `json:"memo"`
		}
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := new(request)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedMeal, err := mh.mealUsecase.Update(id, req.Memo, req.Type, req.Carbs, req.Fat, req.Protein, req.Calories)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := response{
		ID:   updatedMeal.ID,
		Type: updatedMeal.Type,
		Memo: updatedMeal.Memo,
	}

	return c.JSON(http.StatusOK, res)
}

// Delete mealを削除するときのハンドラー
func (mh *mealHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = mh.mealUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
