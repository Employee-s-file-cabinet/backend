package handlers

import (
	"log/slog"

	"github.com/Employee-s-file-cabinet/backend/internal/server/internal/api"
)

var _ api.ServerInterface = (*handler)(nil)

type handler struct {
	fileRepository S3FileRepository
	userRepository UserRepository
	logger         *slog.Logger
}

func New(userRepository UserRepository, s3FileRepository S3FileRepository, logger *slog.Logger) *handler {
	logger = logger.With(slog.String("from", "handler"))

	h := &handler{
		fileRepository: s3FileRepository,
		userRepository: userRepository,
		logger:         logger,
	}

	return h
}
