package usecase

import (
	"github.com/kunikida123456/McdMenuApp/domain/model"
	"github.com/kunikida123456/McdMenuApp/domain/repository"
)

// KindUsecase kind usecaseのinterface
type KindUsecase interface {
	InsertInitialData() (kinds []model.Kind)
}

type kindUsecase struct {
	kindRepo repository.KindRepository
}

func (m *kindUsecase) InsertInitialData(kinds []model.Kind) error {
	_, err := m.kindRepo.Create(kinds)
	if err != nil {
		return nil, err
	}
	return nil
}
