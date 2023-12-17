package postgresql

import (
	"context"

	"github.com/Employee-s-file-cabinet/backend/internal/config"
	"github.com/Employee-s-file-cabinet/backend/internal/model"
	"github.com/Employee-s-file-cabinet/backend/pkg/e"
)

type userStorage struct {
	*model.DB
}

func NewUserStorage(cfg config.PG) (*userStorage, error) {
	const op = "create user storage"

	db, err := model.New(cfg.DSN,
		model.MaxOpenConn(cfg.MaxOpen),
		model.ConnAttempts(cfg.ConnAttempts))
	if err != nil {
		return nil, e.Wrap(op, err)
	}

	return &userStorage{db}, nil
}

func (s *userStorage) ExistUser(ctx context.Context, userID int) (bool, error) {
	const op = "postrgresql user storage: exist user"

	panic("not implemented")
}
