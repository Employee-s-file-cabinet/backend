package user

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Employee-s-file-cabinet/backend/internal/repo/s3"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
	"github.com/Employee-s-file-cabinet/backend/pkg/repoerr"
)

const (
	MaxPhotoSize  = 20 << 20 // bytes
	photoFileName = "photo"
)

func (s *service) Get(ctx context.Context, userID uint64) (*model.User, error) {
	const op = "user service: get user"

	u, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, repoerr.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return u, nil
}

func (s *service) List(ctx context.Context, params model.ListUsersParams) ([]model.User, int, error) {
	const op = "user service: list users"

	users, count, err := s.userRepository.List(ctx, params)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	return users, count, nil
}

func (s *service) DownloadPhoto(ctx context.Context, userID uint64) (model.File, func() error, error) {
	const op = "user service: download photo"

	// if exist, err := s.userRepository.Exist(ctx, userID); err != nil {
	// 	return model.File{}, nil, fmt.Errorf("%s: %w", op, err)
	// } else if !exist {
	// 	return model.File{}, nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
	// }

	f, closeFn, err := s.fileRepository.Download(ctx, strconv.FormatUint(userID, 10), photoFileName)
	if err != nil {
		if errors.Is(err, repoerr.ErrRecordNotFound) {
			return model.File{}, nil, fmt.Errorf("%s: %w", op, ErrPhotoNotFound)
		}
		return model.File{}, nil, fmt.Errorf("%s: %w", op, err)
	}

	return model.File{
		ContentType: f.ContentType,
		Size:        f.Size,
		Reader:      f.Reader,
	}, closeFn, nil
}

func (s *service) UploadPhoto(ctx context.Context, userID uint64, f model.File) error {
	const op = "user service: upload photo"

	if f.Size > MaxPhotoSize {
		return fmt.Errorf("%s: %w", op, ErrPhotoFileSizeTooLarge)
	}

	if exist, err := s.userRepository.Exist(ctx, userID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	} else if !exist {
		return fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	if err := s.fileRepository.Upload(ctx, s3.File{
		Prefix:      strconv.FormatUint(userID, 10),
		Name:        photoFileName,
		Reader:      f.Reader,
		Size:        f.Size,
		ContentType: f.ContentType,
	}); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
