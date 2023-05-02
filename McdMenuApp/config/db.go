package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/kunikida123456/McdMenuApp/domain/model"
	"github.com/subosito/gotenv"
)

func NewDB() *gorm.DB {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := "localhost"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Item{}, model.Kind{})
	db.Model(model.Item{}).AddForeignKey("kind_id", "kind(kind_id)", "RESTRICT", "RESTRICT")

	return db
}
