package mysql

import (
	"minilib/rent/internal/repository"
	"minilib/rent/pkg/model"

	"gorm.io/gorm"
)

type RentRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.RentRepository {
	return &RentRepositoryImpl{
		db: db,
	}
}

func (repo *RentRepositoryImpl) GetAll() ([]model.Rent, error) {
	//TODO: get from logged in user
	var s []model.Rent

	err := repo.db.Find(&s).Error

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (repo *RentRepositoryImpl) Create(rentInput model.RentInput) (model.Rent, error) {
	var createdRent model.Rent = model.Rent{
		UserID:    rentInput.UserID,
		BookID:    rentInput.BookID,
		BookTitle: rentInput.BookTitle,
	}

	result := repo.db.Create(&createdRent)

	if err := result.Error; err != nil {
		return model.Rent{}, err
	}

	err := repo.db.Last(&createdRent).Error

	if err != nil {
		return model.Rent{}, err
	}

	return createdRent, nil
}
