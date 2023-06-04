package rent

import (
	"rent-service/internal/repository"
	"rent-service/pkg/model"
)

type Service struct {
	repo repository.RentRepository
}

func New(repo repository.RentRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (srv *Service) GetAll(userId string) ([]model.Rent, error) {
	return srv.repo.GetAll(userId)
}

func (srv *Service) Create(rentInput model.RentInput) (model.Rent, error) {
	return srv.repo.Create(rentInput)
}
