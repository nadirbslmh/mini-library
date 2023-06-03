package repository

import "minilib/auth/pkg/model"

type AuthRepository interface {
	Register(userInput model.UserInput) (model.User, error)
	Login(userInput model.UserInput) (string, error)
}
