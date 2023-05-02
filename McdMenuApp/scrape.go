package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
			fmt.Println(row)
			data = append(data, row)
		})

		// データを構造体に変換
		items := make([]Item, 0)
		for _, row := range data {
			item := Item{}
			item.Genre = row[0]
			item.Name = row[1]
			if len(row) >= 5 {
				item.Energy, _ = strconv.ParseFloat(row[2], 64)
				item.Protein, _ = strconv.ParseFloat(row[3], 64)
				item.Fat, _ = strconv.ParseFloat(row[4], 64)
				item.Carbohydrates, _ = strconv.ParseFloat(row[5], 64)
			}
			items = append(items, item)
		}

		// データを出力
		fmt.Println(items)
	})
}
