package usecase

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kunikida123456/McdMenuApp/domain/model"
	"github.com/kunikida123456/McdMenuApp/domain/repository"
)

// ItemUsecase item usecaseのinterface
type ItemUsecase interface {
	ScrapeItems() ([]model.Item, error)
	InsertInitialData(items []model.Item)
}

type menuitemUsecase struct {
	menuitemRepo repository.ItemRepository
}

// NewItemUsecase item usecaseのコンストラクタ
func NewItemUsecase(ItemRepo repository.ItemRepository) ItemUsecase {
	return &menuitemUsecase{menuitemRepo: ItemRepo}
}

func (m *menuitemUsecase) InsertInitialData(items []model.Item) error {
	for _, item := range items {
		_, err := m.menuitemRepo.Create(item)
		if err != nil {
			return nil, err
		}
	}
	return nil
}

// FindByID itemをIDで取得するときのユースケース
func (mu *menuitemUsecase) FindByKindID(id int) (*model.Item, error) {
	foundItem, err := mu.menuitemRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundItem, nil
}

// ScrapeItems マックの公式からメニューのテーブルをスクレイピング
func ScrapeData() (Items []model.Item) {
	items := make([]model.Item, 0)
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
		for _, row := range data {
			item := model.Item
			item.KindID = getKindID(row[0])
			item.Name = row[1]
			if len(row) >= 5 {
				item.Energy, _ = strconv.ParseFloat(row[2], 64)
				item.Protein, _ = strconv.ParseFloat(row[3], 64)
				item.Fat, _ = strconv.ParseFloat(row[4], 64)
				item.Carbohydrates, _ = strconv.ParseFloat(row[5], 64)
			}
			items = append(items, item)
		}
	})
	return items

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

func makeRandomMenu()
