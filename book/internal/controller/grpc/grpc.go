package grpc

import (
	"book-service/internal/service/book"
	"book-service/pkg/model"
	"book-service/pkg/util"
	"context"
	"pkg-service/proto_gen"
)

type Server struct {
	proto_gen.UnimplementedBookServiceServer
	service *book.BookService
}

func New(service *book.BookService) *Server {
	return &Server{
		service: service,
	}
}

func (ctrl *Server) GetAllBooks(request *proto_gen.GetAllBooksRequest, stream proto_gen.BookService_GetAllBooksServer) error {
	books, err := ctrl.service.GetAll()

	if err != nil {
		return err
	}

	for _, book := range books {
		stream.Send(&proto_gen.GetAllBooksResponse{
			Book: util.MapToBookPb(book),
		})
	}

	return nil
}

func (ctrl *Server) GetBookByID(ctx context.Context, request *proto_gen.GetBookByIDRequest) (*proto_gen.GetBookByIDResponse, error) {
	var bookId string = request.GetId()

	book, err := ctrl.service.GetByID(bookId)

	if err != nil {
		return &proto_gen.GetBookByIDResponse{
			Status:  "failed",
			Message: "book not found",
		}, err
	}

	return &proto_gen.GetBookByIDResponse{
		Status:  "success",
		Message: "book found",
		Book:    util.MapToBookPb(book),
	}, nil
}

func (ctrl *Server) CreateBook(ctx context.Context, request *proto_gen.CreateBookRequest) (*proto_gen.CreateBookResponse, error) {
	var bookInput model.BookInput = model.BookInput{
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Author:      request.GetAuthor(),
	}

	book, err := ctrl.service.Create(bookInput)

	if err != nil {
		return &proto_gen.CreateBookResponse{
			Status:  "failed",
			Message: "failed to create a book",
		}, err
	}

	return &proto_gen.CreateBookResponse{
		Status:  "success",
		Message: "book created",
		Book:    util.MapToBookPb(book),
	}, nil
}
