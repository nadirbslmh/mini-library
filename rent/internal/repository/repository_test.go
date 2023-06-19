package repository_test

import (
	"errors"
	"os"
	_rentMock "rent-service/internal/mocks/repomocks"
	"rent-service/internal/service/rent"
	"rent-service/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rentRepository _rentMock.RentRepository
var rentService rent.Service

var rentModel model.Rent
var input model.RentInput

func TestMain(m *testing.M) {
	rentModel = model.Rent{
		UserID:    1,
		BookID:    1,
		BookTitle: "test",
	}

	input = model.RentInput{
		UserID:    1,
		BookID:    1,
		BookTitle: "test",
	}

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		rentRepository.On("GetAll", "1").Return([]model.Rent{rentModel}, nil).Once()

		rentService = *rent.New(&rentRepository)

		result, err := rentService.GetAll("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		rentRepository.On("GetAll", "").Return([]model.Rent{}, errors.New("failed")).Once()

		rentService = *rent.New(&rentRepository)

		result, err := rentService.GetAll("")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		rentRepository.On("Create", input).Return(rentModel, nil).Once()

		rentService = *rent.New(&rentRepository)

		result, err := rentService.Create(input)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		rentRepository.On("Create", model.RentInput{}).Return(model.Rent{}, errors.New("failed")).Once()

		rentService = *rent.New(&rentRepository)

		result, err := rentService.Create(model.RentInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
