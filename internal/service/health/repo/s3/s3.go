package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
)

type storage struct {
	minioClient *minio.Client
}

func New(client *minio.Client) (*storage, error) {
	return &storage{
		minioClient: client,
	}, nil
}

func (s *storage) Check(ctx context.Context) error {
	hcancel, err := s.minioClient.HealthCheck(1 * time.Second)
	defer hcancel()

	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	if s.minioClient.IsOffline() {
		return fmt.Errorf("minio client is offline")
	}

	return nil
}
