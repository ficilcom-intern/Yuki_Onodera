package repository

import (
	"error"

	"github.com/kunikida123456/McdMenuApp/domain/model"
)

// Menu_itemRepository kind repositoryのinterface
type KindRepository interface {
	Create(kind *model.Kind) (*model.Kind error)
	FindByKindID(id int) (*model.Kind, error)
}
