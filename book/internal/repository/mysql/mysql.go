package mysql

import (
	"book-service/internal/repository"
	"book-service/pkg/model"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (repo *BookRepositoryImpl) GetAll() ([]model.Book, error) {
	var books []model.Book

	err := repo.db.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *BookRepositoryImpl) GetByID(id string) (model.Book, error) {
	var book model.Book

	err := repo.db.First(&book, "id = ?", id).Error

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (repo *BookRepositoryImpl) Create(bookInput model.BookInput) (model.Book, error) {
	var createdBook model.Book = model.Book{
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Author:      bookInput.Author,
	}

	result := repo.db.Create(&createdBook)

	if err := result.Error; err != nil {
		return model.Book{}, err
	}

	err := repo.db.Last(&createdBook).Error

	if err != nil {
		return model.Book{}, err
	}

	return createdBook, nil
}
