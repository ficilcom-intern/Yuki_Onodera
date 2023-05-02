package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/subosito/gotenv"

	_ "github.com/lib/pq"
)

type Item struct {
	Kind          string
	Name          string
	Energy        float64
	Protein       float64
	Fat           float64
	Carbohydrates float64
}

func main() {
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
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	// スクレイピングして、ヘッダーとデータを取得
	doc, err := goquery.NewDocument("https://www.mcdonalds.co.jp/quality/allergy_Nutrition/nutrient/")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table.allergy-info__table").Each(func(i int, s *goquery.Selection) {
		data := make([][]string, 0)
		s.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
			kind := s.AttrOr("data-kind", "default")
			row := make([]string, 0)
			row = append(row, kind)
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				text := strings.TrimSpace(s.Text())
				if i == 0 || (i >= 2 && i <= 5) {
					row = append(row, text)
				}
			})
			data = append(data, row)
		})

		// データを構造体に変換
		items := make([]Item, 0)
		for _, row := range data {
			item := Item{}
			item.Kind = row[0]
			item.Name = row[1]
			if len(row) >= 5 {
				item.Energy, _ = strconv.ParseFloat(row[2], 64)
				item.Protein, _ = strconv.ParseFloat(row[3], 64)
				item.Fat, _ = strconv.ParseFloat(row[4], 64)
				item.Carbohydrates, _ = strconv.ParseFloat(row[5], 64)
			}
			items = append(items, item)
		}
		for _, item := range items {
			_, err := db.Exec(
				"INSERT INTO item (kind_id, name, energy, protein, fat, carbohydrates) VALUES ($1, $2, $3, $4, $5, $6)",
				getKindID(item.Kind), item.Name, item.Energy, item.Protein, item.Fat, item.Carbohydrates)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func getKindID(s string) int {
	switch s {
	case "ドリンク":
		return 1
	case "バーガー":
		return 2
	case "サイド":
		return 3
	case "バリスタ":
		return 4
	default:
		return 0
	}
}
