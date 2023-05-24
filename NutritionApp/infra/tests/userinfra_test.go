package infra_test

import (
	"regexp"
	"testing"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/infra"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestUserRepository_CreateUser(t *testing.T) {
	mockDB, mock, err := GetNewDbMock()
	assert.NoError(t, err)

	mock.ExpectBegin() // トランザクションの開始を期待

	userRepo := infra.NewUserRepository(mockDB)

	user := &model.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("name","email","password","created_at") VALUES ($1,$2,$3,$4)`)).
		WithArgs(user.Name, user.Email, user.Password, sqlmock.AnyArg()).
		WillReturnRows(rows)

	// トランザクションのコミットを期待
	mock.ExpectCommit()

	// CreateUserメソッドの呼び出し
	result, err := userRepo.CreateUser(user)

	// エラーがないことを検証
	assert.NoError(t, err)

	// 返されたユーザーが正しいことを検証
	assert.Equal(t, user, result)

	// モックの期待が正しいことを検証
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	mockDB, mock, err := GetNewDbMock()
	assert.NoError(t, err)

	userRepo := infra.NewUserRepository(mockDB)

	email := "john@example.com"
	user := &model.User{
		Name:     "John Doe",
		Email:    email,
		Password: "password",
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"."id" LIMIT 1`).
		WithArgs("john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(0, "John Doe", "john@example.com", "password"))

	result, err := userRepo.GetUserByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, user, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetUserByEmail_NotFound(t *testing.T) {
	mockDB, mock, err := GetNewDbMock()
	assert.NoError(t, err)

	userRepo := infra.NewUserRepository(mockDB)

	email := "john@example.com"

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"."id" LIMIT 1`).
		WithArgs("john@example.com").
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := userRepo.GetUserByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, int(0), result.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
