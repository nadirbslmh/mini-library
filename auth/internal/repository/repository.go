package repository

import "auth-service/pkg/model"

type AuthRepository interface {
	Register(userInput model.UserInput) (model.User, error)
	Login(userInput model.UserInput) (string, error)
}
