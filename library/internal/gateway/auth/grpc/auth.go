package grpc

import (
	authmodel "auth-service/pkg/model"
	"context"
	"log"
	"pkg-service/discovery"
	"pkg-service/model"

	"pkg-service/proto_gen"
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

	var client proto_gen.AuthServiceClient = proto_gen.NewAuthServiceClient(conn)

	request := &proto_gen.RegisterRequest{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	res, err := client.Register(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[authmodel.User]{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
		Data: authmodel.User{
			ID:        uint(res.User.GetId()),
			CreatedAt: res.User.GetCreatedAt().AsTime(),
			UpdatedAt: res.User.GetUpdatedAt().AsTime(),
			DeletedAt: gorm.DeletedAt{
				Time:  res.User.GetDeletedAt().GetTime().AsTime(),
				Valid: res.User.GetDeletedAt().GetValid(),
			},
			Email:    res.User.GetEmail(),
			Password: res.User.GetPassword(),
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

	var client proto_gen.AuthServiceClient = proto_gen.NewAuthServiceClient(conn)

	request := &proto_gen.LoginRequest{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	res, err := client.Login(ctx, request)

	if err != nil {
		return nil, err
	}

	return &model.Response[string]{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
		Data:    res.GetToken(),
	}, nil
}
