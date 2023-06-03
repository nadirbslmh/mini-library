package rent

import (
	"minilib/rent/internal/repository"
	"minilib/rent/pkg/model"
)

type RentService struct {
	repo repository.RentRepository
}

func New(repo repository.RentRepository) *RentService {
	return &RentService{
		repo: repo,
	}
}

func (srv *RentService) GetAll() ([]model.Rent, error) {
	return srv.repo.GetAll()
}

func (srv *RentService) Create(rentInput model.RentInput) (model.Rent, error) {
	return srv.repo.Create(rentInput)
}
