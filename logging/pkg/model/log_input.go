package model

import "github.com/go-playground/validator/v10"

type LogInput struct {
	UserID    int    `json:"user_id" validate:"required"`
	BookID    int    `json:"book_id" validate:"required"`
	BookTitle string `json:"book_title" validate:"required"`
}

func (input *LogInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
