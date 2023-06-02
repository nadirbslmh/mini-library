package library

import (
	"context"
	bookmodel "minilib/book/pkg/model"
	"minilib/library/internal/gateway/book/http"
	"minilib/library/pkg/model"
)

type Service struct {
	gateway http.Gateway
}

func New() *Service {
	return &Service{
		gateway: http.Gateway{},
	}
}

func (srv *Service) GetAll(ctx context.Context) (*model.Response, error) {
	data, err := srv.gateway.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *Service) Create(ctx context.Context, bookInput bookmodel.Book) (*model.Response, error) {
	data, err := srv.gateway.Create(ctx, bookInput)

	if err != nil {
		return nil, err
	}

	return data, nil
}
