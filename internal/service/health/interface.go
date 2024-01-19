package health

import "context"

type dbRepository interface {
	Check(ctx context.Context) error
}

type fileRepository interface {
	Check(ctx context.Context) error
}
