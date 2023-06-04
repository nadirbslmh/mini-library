package auth

import (
	"auth-service/internal/repository"
	"auth-service/pkg/model"
)

type AuthService struct {
	repo repository.AuthRepository
}

func New(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (srv *AuthService) Register(userInput model.UserInput) (model.User, error) {
	return srv.repo.Register(userInput)
}

func (srv *AuthService) Login(userInput model.UserInput) (string, error) {
	return srv.repo.Login(userInput)
}
