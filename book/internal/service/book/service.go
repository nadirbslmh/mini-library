package book

import (
	"minilib/book/internal/repository/memory"
	"minilib/book/pkg/model"
)

type Service struct {
	repo *memory.Repository
}

func New() *Service {
	return &Service{
		repo: &memory.Repository{},
	}
}

func (srv *Service) GetAll() []model.Book {
	return srv.repo.GetAll()
}

func (srv *Service) Create(book model.Book) model.Book {
	return srv.repo.Create(book)
}
