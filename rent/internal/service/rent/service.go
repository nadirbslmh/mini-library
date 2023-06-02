package rent

import (
	"minilib/rent/internal/repository/memory"
	"minilib/rent/pkg/model"
)

type Service struct {
	repo *memory.Repository
}

func New() *Service {
	return &Service{
		repo: &memory.Repository{},
	}
}

func (srv *Service) GetAll() []model.Rent {
	return srv.repo.GetAll()
}

func (srv *Service) Create(book model.Rent) model.Rent {
	return srv.repo.Create(book)
}
