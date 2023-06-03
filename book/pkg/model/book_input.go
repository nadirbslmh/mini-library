package model

import "github.com/go-playground/validator/v10"

type BookInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
}

func (input *BookInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
