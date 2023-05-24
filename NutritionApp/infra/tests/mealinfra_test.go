package infra_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/infra"
)

func TestMealRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// モックデータ
	meal := &model.Meal{
		UserID:   1,
		Memo:     "memo",
		MealType: "Breakfast",
		Carbs:    10,
		Fat:      20,
		Protein:  30,
		Calories: 250,
	}

	// MealRepositoryの生成
	repo := infra.NewMealRepository(gormDB)

	// Createメソッドのテスト
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "meal" (.+) RETURNING "memo","meal_type","carbs","fat","protein","calories","id"`).
		WithArgs(meal.UserID, meal.Memo, meal.MealType, sqlmock.AnyArg(), sqlmock.AnyArg(), meal.Carbs, meal.Fat, meal.Protein, meal.Calories).
		WillReturnRows(sqlmock.NewRows([]string{"memo", "meal_type", "carbs", "fat", "protein", "calories", "id"}).
			AddRow(meal.Memo, meal.MealType, meal.Carbs, meal.Fat, meal.Protein, meal.Calories, 1))
	mock.ExpectCommit()

	createdMeal, err := repo.Create(meal)
	assert.NoError(t, err)
	assert.Equal(t, meal, createdMeal)

	// Updateメソッドのテスト
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "meal" (.+)`).
		WithArgs(meal.UserID, meal.Memo, meal.MealType, meal.Carbs, meal.Fat, meal.Protein, meal.Calories, meal.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	updatedMeal, err := repo.Update(meal)
	assert.NoError(t, err)
	assert.Equal(t, meal, updatedMeal)

	// Deleteメソッドのテスト
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "meal" WHERE (.+)`).
		WithArgs(meal.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.Delete(meal)
	assert.NoError(t, err)

	// FindByIDメソッドのテスト
	mock.ExpectQuery(`SELECT * FROM "meal" WHERE (.+)`).
		WithArgs(meal.ID, meal.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "memo", "meal_type", "carbs", "fat", "protein", "calories"}).
			AddRow(meal.ID, meal.UserID, meal.Memo, meal.MealType, meal.Carbs, meal.Fat, meal.Protein, meal.Calories))

	foundMeal, err := repo.FindByID(meal.ID, meal.UserID)
	assert.NoError(t, err)
	assert.Equal(t, meal, foundMeal)

	// FindAllメソッドのテスト
	mock.ExpectQuery(`SELECT * FROM "meal" WHERE (.+)`).
		WithArgs(meal.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "memo", "meal_type", "carbs", "fat", "protein", "calories"}).
			AddRow(meal.ID, meal.UserID, meal.Memo, meal.MealType, meal.Carbs, meal.Fat, meal.Protein, meal.Calories))

	allMeals, err := repo.FindAll(meal.UserID)
	assert.NoError(t, err)
	assert.Len(t, allMeals, 1)
	assert.Equal(t, meal, allMeals[0])

	// 予期されたクエリが実行されたか検証
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
