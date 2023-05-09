package config

import (
	"fmt"
	"log"
	"os"

	"kunikida123456/NutritionApp/domain/model"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := "go_db"
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.User{}, model.Meal{})
	// db.Model(model.Item{}).AddForeignKey("id", "kind(id)", "RESTRICT", "RESTRICT")
}