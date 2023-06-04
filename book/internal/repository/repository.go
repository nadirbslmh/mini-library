package repository

import "book-service/pkg/model"

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetByID(id string) (model.Book, error)
	Create(bookInput model.BookInput) (model.Book, error)
}
