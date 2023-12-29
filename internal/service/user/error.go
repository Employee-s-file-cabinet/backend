package user

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrPhotoNotFound         = errors.New("photo not found")
	ErrPhotoFileSizeTooLarge = errors.New("photo file size too large")
)
