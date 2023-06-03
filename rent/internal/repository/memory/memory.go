package memory

import "minilib/rent/pkg/model"

type Repository struct{}

var database []model.Rent = []model.Rent{}

func (r *Repository) GetAll() []model.Rent {
	return database
}

func (r *Repository) Create(rent model.Rent) model.Rent {
	database = append(database, rent)
	return rent
}
