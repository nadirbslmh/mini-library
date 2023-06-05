package mysql

import (
	"rent-service/internal/repository"
	"rent-service/pkg/model"

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

func (repo *RentRepositoryImpl) GetAll(userId string) ([]model.Rent, error) {
	var rents []model.Rent

	err := repo.db.Where("user_id = ?", userId).Find(&rents).Error

	if err != nil {
		return nil, err
	}

	return rents, nil
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
