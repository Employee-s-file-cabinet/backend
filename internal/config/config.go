package config

import (
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http"
	repopg "github.com/Employee-s-file-cabinet/backend/internal/repo/postgresql"
	repos3 "github.com/Employee-s-file-cabinet/backend/internal/repo/s3"
)

type Config struct {
	LogLevel string        `env:"LOG_LEVEL" env-default:"debug"`
	HTTP     http.Config   `env-prefix:"HTTP_"`
	PG       repopg.Config `env-prefix:"PG_"`
	S3       repos3.Config `env-prefix:"S3_"`
}

// New создаёт объект Config.
func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
