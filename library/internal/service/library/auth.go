package library

import (
	authmodel "auth-service/pkg/model"
	"context"
	"library-service/internal/gateway/auth/grpc"
	"pkg-service/model"
)

type AuthService struct {
	gateway grpc.Gateway
}

func NewAuthService(gateway grpc.Gateway) *AuthService {
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
