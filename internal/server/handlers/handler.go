package handlers

import (
	"log/slog"

	"github.com/Employee-s-file-cabinet/backend/internal/server/internal/api"
)

var _ api.ServerInterface = (*handler)(nil)

type handler struct {
	fileUploader   S3FileUploader
	userRepository UserRepository
	logger         *slog.Logger
}

func New(userRepository UserRepository, s3FileUploader S3FileUploader, logger *slog.Logger) *handler {
	logger = logger.With(slog.String("from", "handler"))

	h := &handler{
		fileUploader:   s3FileUploader,
		userRepository: userRepository,
		logger:         logger,
	}

	return h
}
