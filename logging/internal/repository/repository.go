package repository

import "logging-service/pkg/model"

type LogRepository interface {
	Write(logInput model.LogInput) (*model.Log, error)
}
