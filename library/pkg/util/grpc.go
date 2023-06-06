package util

import (
	"book-service/pkg/model"
	"pkg-service/proto_gen"

	"gorm.io/gorm"
)

func MapToBookModel(book *proto_gen.Book) model.Book {
	return model.Book{
		ID:        uint(book.GetId()),
		CreatedAt: book.GetCreatedAt().AsTime(),
		UpdatedAt: book.GetUpdatedAt().AsTime(),
		DeletedAt: gorm.DeletedAt{
			Time:  book.GetDeletedAt().Time.AsTime(),
			Valid: book.GetDeletedAt().GetValid(),
		},
		Title:       book.GetTitle(),
		Description: book.GetDescription(),
		Author:      book.GetAuthor(),
	}
}
