package repository_test

import (
	_authMock "auth-service/internal/mocks/repomocks"
	"auth-service/internal/service/auth"
	"auth-service/pkg/model"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var authRepository _authMock.AuthRepository
var authService auth.AuthService

var user model.User
var input model.UserInput

func TestMain(m *testing.M) {
	user = model.User{
		Email:    "test@test.com",
		Password: "123123",
	}

	input = model.UserInput{
		Email:    "test@test.com",
		Password: "123123",
	}

	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		authRepository.On("Register", input).Return(user, nil).Once()

		authService = *auth.New(&authRepository)

		result, err := authService.Register(input)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		authRepository.On("Register", model.UserInput{}).Return(model.User{}, errors.New("failed")).Once()

		authService = *auth.New(&authRepository)

		result, err := authService.Register(model.UserInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		authRepository.On("Login", input).Return("token", nil).Once()

		authService = *auth.New(&authRepository)

		result, err := authService.Login(input)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Login | Invalid", func(t *testing.T) {
		authRepository.On("Login", model.UserInput{}).Return("", errors.New("failed")).Once()

		authService = *auth.New(&authRepository)

		result, err := authService.Login(model.UserInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
