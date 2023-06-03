package library

import (
	"context"
	"minilib/library/internal/gateway/rent/http"
	"minilib/pkg/model"
	rentmodel "minilib/rent/pkg/model"
)

type RentService struct {
	gateway http.Gateway
}

func NewRentService(gateway http.Gateway) *RentService {
	return &RentService{
		gateway: gateway,
	}
}

func (srv *RentService) GetAllRents(ctx context.Context) (*model.Response[[]rentmodel.Rent], error) {
	data, err := srv.gateway.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *RentService) CreateRent(ctx context.Context, bookInput rentmodel.RentInput) (*model.Response[rentmodel.Rent], error) {
	data, err := srv.gateway.Create(ctx, bookInput)

	if err != nil {
		return nil, err
	}

	return data, nil
}
