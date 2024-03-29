package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/Employee-s-file-cabinet/backend/internal/config/env"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/smtp"
	repopg "github.com/Employee-s-file-cabinet/backend/internal/repo/postgresql"
	repos3 "github.com/Employee-s-file-cabinet/backend/internal/repo/s3"
	"github.com/Employee-s-file-cabinet/backend/internal/service/recovery"
)

type Config struct {
	EnvType  env.Type        `env:"ENV_TYPE" env-required:"production"`
	LogLevel slog.Level      `env:"LOG_LEVEL" env-default:"INFO" env-description:"importance or severity of a log event (DEBUG/INFO/WARN/ERROR)"`
	Recovery recovery.Config `env-prefix:"RECOVERY_"`
	HTTP     http.Config     `env-prefix:"HTTP_"`
	PG       repopg.Config   `env-prefix:"PG_"`
	S3       repos3.Config   `env-prefix:"S3_"`
	Mail     smtp.Config     `env-prefix:"MAIL_"`
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
