package memory

import "minilib/book/pkg/model"

type Repository struct{}

var database []model.Book = []model.Book{}

func (r *Repository) GetAll() []model.Book {
	return database
}

func (r *Repository) Create(book model.Book) model.Book {
	database = append(database, book)
	return book
}
