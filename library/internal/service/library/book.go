package library

import (
	"context"
	bookmodel "minilib/book/pkg/model"
	"minilib/library/internal/gateway/book/http"
	"minilib/pkg/model"
)

type BookService struct {
	gateway http.Gateway
}

func NewBookService(gateway http.Gateway) *BookService {
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
