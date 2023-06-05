package grpc

import (
	"auth-service/internal/service/auth"
	"auth-service/pkg/model"
	"context"
	"pkg-service/auth_gen"
)

type Server struct {
	auth_gen.UnimplementedAuthServiceServer
	service *auth.AuthService
}

func New(service *auth.AuthService) *Server {
	return &Server{
		service: service,
	}
}

func (ctrl *Server) Register(ctx context.Context, request *auth_gen.RegisterRequest) (*auth_gen.RegisterResponse, error) {
	userInput := model.UserInput{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	user, err := ctrl.service.Register(userInput)

	if err != nil {
		return &auth_gen.RegisterResponse{
			Status:  "failed",
			Message: "failed to create a user",
		}, err
	}

	return &auth_gen.RegisterResponse{
		Status:  "success",
		Message: "user created",
		User: &auth_gen.User{
			Id:       uint32(user.ID),
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func (ctrl *Server) Login(ctx context.Context, request *auth_gen.LoginRequest) (*auth_gen.LoginResponse, error) {
	userInput := model.UserInput{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	token, err := ctrl.service.Login(userInput)

	if err != nil {
		return &auth_gen.LoginResponse{
			Status:  "failed",
			Message: "login failed. invalid username or password",
		}, err
	}

	return &auth_gen.LoginResponse{
		Status:  "success",
		Message: "login success",
		Token:   token,
	}, nil
}
