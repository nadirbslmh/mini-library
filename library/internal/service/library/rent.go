package library

import (
	"context"
	"minilib/library/internal/gateway/rent/http"
	"minilib/pkg/model"
	rentmodel "rent-service/pkg/model"
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
	return srv.gateway.GetAll(ctx)
}

func (srv *RentService) CreateRent(ctx context.Context, bookInput rentmodel.RentInput) (*model.Response[rentmodel.Rent], error) {
	return srv.gateway.Create(ctx, bookInput)
}
