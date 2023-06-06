package grpc

import (
	"book-service/internal/service/book"
	"book-service/pkg/model"
	"context"
	"pkg-service/book_gen"
)

type Server struct {
	book_gen.UnimplementedBookServiceServer
	service *book.BookService
}

func New(service *book.BookService) *Server {
	return &Server{
		service: service,
	}
}

func (ctrl *Server) GetAllBooks(request *book_gen.GetAllBooksRequest, stream book_gen.BookService_GetAllBooksServer) error {
	books, err := ctrl.service.GetAll()

	if err != nil {
		return err
	}

	for _, book := range books {
		stream.Send(&book_gen.GetAllBooksResponse{
			Book: &book_gen.Book{
				Id:          uint32(book.ID),
				Title:       book.Title,
				Description: book.Description,
				Author:      book.Author,
			},
		})
	}

	return nil
}

func (ctrl *Server) GetBookByID(ctx context.Context, request *book_gen.GetBookByIDRequest) (*book_gen.GetBookByIDResponse, error) {
	var bookId string = request.GetId()

	book, err := ctrl.service.GetByID(bookId)

	if err != nil {
		return &book_gen.GetBookByIDResponse{
			Status:  "failed",
			Message: "book not found",
		}, err
	}

	return &book_gen.GetBookByIDResponse{
		Status:  "success",
		Message: "book found",
		Book: &book_gen.Book{
			Id:          uint32(book.ID),
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		},
	}, nil
}

func (ctrl *Server) CreateBook(ctx context.Context, request *book_gen.CreateBookRequest) (*book_gen.CreateBookResponse, error) {
	var bookInput model.BookInput = model.BookInput{
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Author:      request.GetAuthor(),
	}

	book, err := ctrl.service.Create(bookInput)

	if err != nil {
		return &book_gen.CreateBookResponse{
			Status:  "failed",
			Message: "failed to create a book",
		}, err
	}

	return &book_gen.CreateBookResponse{
		Status:  "success",
		Message: "book created",
		Book: &book_gen.Book{
			Id:          uint32(book.ID),
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		},
	}, nil
}
