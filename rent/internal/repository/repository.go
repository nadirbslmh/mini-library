package repository

import "rent-service/pkg/model"

type RentRepository interface {
	GetAll(userId string) ([]model.Rent, error)
	Create(rentInput model.RentInput) (model.Rent, error)
}
