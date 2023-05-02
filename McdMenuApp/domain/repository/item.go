package repository

import (
	"github.com/kunikida123456/McdMenuApp/domain/model"
)

type ItemRepository interface {
	Create(item *model.Item) (*model.Item, error)
	FindByKindID(id int) (*model.Item, error)
}
