package user

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrEducationNotFound      = errors.New("education not found")
	ErrTrainingNotFound       = errors.New("training not found")
	ErrPassportNotFound       = errors.New("passport not found")
	ErrVisaNotFound           = errors.New("visa not found")
	ErrVacationNotFound       = errors.New("vacation not found")
	ErrUserOrPassportNotFound = errors.New("user or passport not found")
	ErrPhotoFileNotFound      = errors.New("photo file not found")
	ErrPhotoFileNotModified   = errors.New("photo file not modified")
	ErrPhotoFileSizeTooLarge  = errors.New("photo file size too large")
	ErrScanFileNotFound       = errors.New("scan file not found")
	ErrScanFileSizeTooLarge   = errors.New("scan file size too large")
)
