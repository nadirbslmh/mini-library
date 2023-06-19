package repository_test

import (
	_bookMock "book-service/internal/mocks/repomocks"
	"book-service/internal/service/book"
	"book-service/pkg/model"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookRepository _bookMock.BookRepository
var bookService book.BookService

var bookModel model.Book
var input model.BookInput

func TestMain(m *testing.M) {
	bookModel = model.Book{
		Title:       "test",
		Description: "test",
		Author:      "test",
	}

	input = model.BookInput{
		Title:       "test",
		Description: "test",
		Author:      "test",
	}

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		bookRepository.On("GetAll").Return([]model.Book{bookModel}, nil).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		bookRepository.On("GetAll").Return([]model.Book{}, errors.New("failed")).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		bookRepository.On("GetByID", "1").Return(bookModel, nil).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		bookRepository.On("GetByID", "").Return(model.Book{}, errors.New("failed")).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.GetByID("")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		bookRepository.On("Create", input).Return(bookModel, nil).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.Create(input)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		bookRepository.On("Create", model.BookInput{}).Return(model.Book{}, errors.New("failed")).Once()

		bookService = *book.New(&bookRepository)

		result, err := bookService.Create(model.BookInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
