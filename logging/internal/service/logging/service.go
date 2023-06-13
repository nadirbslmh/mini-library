package logging

import (
	"logging-service/internal/repository"
	"logging-service/pkg/model"
)

type LogService struct {
	repo repository.LogRepository
}

func New(repo repository.LogRepository) *LogService {
	return &LogService{
		repo: repo,
	}
}

func (srv *LogService) Write(logInput model.LogInput) (*model.Log, error) {
	return srv.repo.Write(logInput)
}
