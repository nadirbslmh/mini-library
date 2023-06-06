package grpc

import (
	authmodel "auth-service/pkg/model"
	"context"
	"log"
	"pkg-service/auth_gen"
	"pkg-service/discovery"
	"pkg-service/model"
	"pkg-service/util"

	"gorm.io/gorm"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{
		registry: registry,
	}
}

func (g *Gateway) Register(ctx context.Context, userInput authmodel.UserInput) (*model.Response[authmodel.User], error) {
	log.Println("calling auth service with gRPC: register")

	conn, err := util.ConnectgRPCService(ctx, "auth", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client auth_gen.AuthServiceClient = auth_gen.NewAuthServiceClient(conn)

	request := &auth_gen.RegisterRequest{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	res, err := client.Register(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[authmodel.User]{
		Status:  res.Status,
		Message: res.Message,
		Data: authmodel.User{
			ID:        uint(res.User.Id),
			CreatedAt: res.User.CreatedAt.AsTime(),
			UpdatedAt: res.User.UpdatedAt.AsTime(),
			DeletedAt: gorm.DeletedAt{
				Time:  res.User.DeletedAt.GetTime().AsTime(),
				Valid: res.User.DeletedAt.Valid,
			},
			Email:    res.User.Email,
			Password: res.User.Password,
		},
	}, nil
}

func (g *Gateway) Login(ctx context.Context, userInput authmodel.UserInput) (*model.Response[string], error) {
	log.Println("calling auth service with gRPC: login")

	conn, err := util.ConnectgRPCService(ctx, "auth", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var client auth_gen.AuthServiceClient = auth_gen.NewAuthServiceClient(conn)

	request := &auth_gen.LoginRequest{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	res, err := client.Login(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[string]{
		Status:  res.Status,
		Message: res.Message,
		Data:    res.Token,
	}, nil
}
