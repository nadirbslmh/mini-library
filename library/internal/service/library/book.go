package library

import (
	bookmodel "book-service/pkg/model"
	"context"
	"library-service/internal/gateway/book/grpc"
	"pkg-service/model"
)

type BookService struct {
	gateway grpc.Gateway
}

func NewBookService(gateway grpc.Gateway) *BookService {
	return &BookService{
		gateway: gateway,
	}
}

func (srv *BookService) GetAllBooks(ctx context.Context) (*model.Response[[]bookmodel.Book], error) {
	return srv.gateway.GetAll(ctx)
}

func (srv *BookService) GetBookByID(ctx context.Context, id string) (*model.Response[bookmodel.Book], error) {
	return srv.gateway.GetByID(ctx, id)
}

func (srv *BookService) CreateBook(ctx context.Context, bookInput bookmodel.BookInput) (*model.Response[bookmodel.Book], error) {
	return srv.gateway.Create(ctx, bookInput)
}
