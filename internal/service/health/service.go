package health

import (
	"context"

	"github.com/Employee-s-file-cabinet/backend/internal/config"
	"github.com/Employee-s-file-cabinet/backend/internal/service/health/model"
)

type service struct {
	appInfo        config.App
	dbRepository   dbRepository
	fileRepository fileRepository
}

func New(info config.App, dbr dbRepository, fr fileRepository) *service {
	return &service{
		appInfo:        info,
		dbRepository:   dbr,
		fileRepository: fr,
	}
}

func (s *service) HealthCheck(ctx context.Context) model.ServiceStatus {
	hs := model.ServiceStatus{
		Version:     s.appInfo.Version,
		Date:        s.appInfo.Date,
		Commit:      s.appInfo.Commit,
		Database:    "OK",
		FileStorage: "OK",
	}

	if err := s.dbRepository.Check(ctx); err != nil {
		hs.Database = err.Error()
	}

	if err := s.fileRepository.Check(ctx); err != nil {
		hs.FileStorage = err.Error()
	}

	return hs
}
