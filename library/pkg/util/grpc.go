package util

import (
	bookmodel "book-service/pkg/model"
	"pkg-service/proto_gen"
	rentmodel "rent-service/pkg/model"

	"gorm.io/gorm"
)

func MapToBookModel(book *proto_gen.Book) bookmodel.Book {
	return bookmodel.Book{
		ID:        uint(book.GetId()),
		CreatedAt: book.GetCreatedAt().AsTime(),
		UpdatedAt: book.GetUpdatedAt().AsTime(),
		DeletedAt: gorm.DeletedAt{
			Time:  book.GetDeletedAt().GetTime().AsTime(),
			Valid: book.GetDeletedAt().GetValid(),
		},
		Title:       book.GetTitle(),
		Description: book.GetDescription(),
		Author:      book.GetAuthor(),
	}
}

func MapToRentModel(rent *proto_gen.Rent) rentmodel.Rent {
	return rentmodel.Rent{
		ID:        uint(rent.GetId()),
		CreatedAt: rent.GetCreatedAt().AsTime(),
		UpdatedAt: rent.GetUpdatedAt().AsTime(),
		DeletedAt: gorm.DeletedAt{
			Time:  rent.GetDeletedAt().GetTime().AsTime(),
			Valid: rent.GetDeletedAt().GetValid(),
		},
		UserID:    int(rent.GetUserId()),
		BookID:    int(rent.GetBookId()),
		BookTitle: rent.GetBookTitle(),
	}
}
