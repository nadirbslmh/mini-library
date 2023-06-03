package library

import (
	"context"
	bookmodel "minilib/book/pkg/model"
	"minilib/library/internal/gateway/book/http"
	"minilib/pkg/model"
)

type Service struct {
	gateway http.Gateway
}

func New(gateway http.Gateway) *Service {
	return &Service{
		gateway: gateway,
	}
}

func (srv *Service) GetAll(ctx context.Context) (*model.Response, error) {
	data, err := srv.gateway.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *Service) GetByID(ctx context.Context, id string) (*model.Response, error) {
	data, err := srv.gateway.GetByID(ctx, id)

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
