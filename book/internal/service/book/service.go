package book

import (
	"minilib/book/internal/repository"
	"minilib/book/pkg/model"
)

type BookService struct {
	repo repository.BookRepository
}

func New(repo repository.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (srv *BookService) GetAll() ([]model.Book, error) {
	return srv.repo.GetAll()
}

func (srv *BookService) GetByID(id string) (model.Book, error) {
	return srv.repo.GetByID(id)
}
func (srv *BookService) Create(bookInput model.BookInput) (model.Book, error) {
	return srv.repo.Create(bookInput)
}
