package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/kunikida123456/McdMenuApp/domain/model"
	"github.com/kunikida123456/McdMenuApp/domain/repository"
)

type KindRepository struct {
	Conn *gorm.DB
}

// NewKindRepository kind repositoryのコンストラクタ
func NewKindRepository(conn *gorm.DB) repository.KindRepository {
	return &KindRepository{Conn: conn}
}

func (kr *KindRepository) Create(kind *model.Kind) (*model.Kind, error) {
	if err := kr.Conn.Create(&kind).Error; err != nil {
		return nil, err
	}
	return kind, nil
}
