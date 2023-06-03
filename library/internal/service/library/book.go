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

func New(gateway http.Gateway) *BookService {
	return &BookService{
		gateway: gateway,
	}
}

func (srv *BookService) GetAllBooks(ctx context.Context) (*model.Response, error) {
	data, err := srv.gateway.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *BookService) GetBookByID(ctx context.Context, id string) (*model.Response, error) {
	data, err := srv.gateway.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *BookService) CreateBook(ctx context.Context, bookInput bookmodel.BookInput) (*model.Response, error) {
	data, err := srv.gateway.Create(ctx, bookInput)

	if err != nil {
		return nil, err
	}

	return data, nil
}
