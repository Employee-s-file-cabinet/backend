package handlers

import (
	"context"
	"log/slog"

	"github.com/Employee-s-file-cabinet/backend/internal/server/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/storage/s3"
)

var _ api.ServerInterface = (*handler)(nil)

type S3FileRepository interface {
	UploadFile(context.Context, s3.File) error
	DownloadFile(ctx context.Context, prefix, name string) (file s3.File, closeFn func() error, err error)
}

type UserRepository interface {
	ExistUser(ctx context.Context, userID int) (bool, error)
}

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
