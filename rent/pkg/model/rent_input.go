package model

import "github.com/go-playground/validator/v10"

type RentInput struct {
	UserID    int    `json:"user_id" validate:"required"`
	BookID    int    `json:"book_id" validate:"required"`
	BookTitle string `json:"book_title" validate:"required"`
}

func (input *RentInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
