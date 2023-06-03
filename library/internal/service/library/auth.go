package library

import (
	"context"
	authmodel "minilib/auth/pkg/model"
	"minilib/library/internal/gateway/auth/http"
	"minilib/pkg/model"
)

type AuthService struct {
	gateway http.Gateway
}

func NewAuthService(gateway http.Gateway) *AuthService {
	return &AuthService{
		gateway: gateway,
	}
}

func (srv *AuthService) Register(ctx context.Context, userInput authmodel.UserInput) (*model.Response[authmodel.User], error) {
	return srv.gateway.Register(ctx, userInput)
}

func (srv *AuthService) Login(ctx context.Context, userInput authmodel.UserInput) (*model.Response[string], error) {
	return srv.gateway.Login(ctx, userInput)
}
