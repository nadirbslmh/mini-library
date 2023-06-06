package util

import (
	"book-service/pkg/model"
	"pkg-service/proto_gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToBookPb(book model.Book) *proto_gen.Book {
	return &proto_gen.Book{
		Id:        uint32(book.ID),
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
		DeletedAt: &proto_gen.DeletedAt{
			Time:  timestamppb.New(book.DeletedAt.Time),
			Valid: book.DeletedAt.Valid,
		},
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}
}
