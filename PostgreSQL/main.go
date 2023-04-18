package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	HOST     = os.Getenv("localhost")
	DATABASE = os.Getenv("postgres")
	USER     = os.Getenv("postgres")
	PASSWORD = os.Getenv("postgres")
)

type User struct {
	Id   int    `db:"user_id"`
	Name string `db:"user_name"`
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	message := os.Getenv("SAMPLE_MESSAGE")

	fmt.Println(message)
}

func main() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE users (user_id serial PRIMARY KEY, user_name VARCHAR(50));")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO
			users (user_name)
		VALUES
			('太郎'),
			('二郎'),
			('三郎')
	`)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Name: %s\n", strconv.Itoa(user.Id), user.Name)
	}
}
