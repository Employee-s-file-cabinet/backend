package postgres

import (
	"context"
	"time"

	pq "github.com/Employee-s-file-cabinet/backend/pkg/postgresql"
)

type storage struct {
	*pq.DB
}

func New(db *pq.DB) (*storage, error) {
	return &storage{
		DB: db,
	}, nil
}

func (s *storage) Check(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.DB.Ping(ctx)
}
