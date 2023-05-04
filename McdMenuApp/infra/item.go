package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/kunikida123456/McdMenuApp/domain/model"
	"github.com/kunikida123456/McdMenuApp/domain/repository"
)

type ItemRepository struct {
	Conn *gorm.DB
}

// NewItemRepository item repositoryのコンストラクタ
func NewItemRepository(conn *gorm.DB) repository.ItemRepository {
	return &ItemRepository{Conn: conn}
}

func (ir *ItemRepository) Create(item *model.Item) (*model.Item, error) {
	if err := ir.Conn.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (ir *ItemRepository) FindByKindID(kind_id int) (*model.Item, error) {
	item := &model.Item{ID: kind_id}
	if err := ir.Conn.First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}
