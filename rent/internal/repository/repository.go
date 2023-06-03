package repository

import "minilib/rent/pkg/model"

type RentRepository interface {
	GetAll() ([]model.Rent, error)
	Create(rentInput model.RentInput) (model.Rent, error)
}
