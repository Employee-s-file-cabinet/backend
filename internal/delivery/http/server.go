package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jub0bs/fcors"

	"github.com/Employee-s-file-cabinet/backend/internal/config/env"
	srvErrors "github.com/Employee-s-file-cabinet/backend/internal/delivery/http/errors"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/handlers"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/middleware"
)

const baseURL = "/api/v1"

const (
	defaultIdleTimeout    = time.Minute
	defaultReadTimeout    = 5 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultShutdownPeriod = 30 * time.Second
)

type server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func New(cfg Config, envType env.Type,
	userService handlers.UserService,
	authService handlers.AuthService,
	logger *slog.Logger) (*server, error) {
	logger = logger.With(slog.String("from", "http-server"))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelWarn),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	s := &server{
		httpServer: srv,
		logger:     logger,
	}

	handler := handlers.New(userService, authService, logger)

	mux := chi.NewRouter()
	mux.NotFound(srvErrors.NotFound)
	mux.MethodNotAllowed(srvErrors.MethodNotAllowed)
	mux.Use(middleware.LogAccess)
	mux.Use(middleware.RecoverPanic)

	switch envType {
	case env.Development, env.Testing:
		cors, err := fcors.AllowAccessWithCredentials(
			fcors.FromOrigins(
				"https://localhost:*",
				"http://localhost:*"),
			fcors.WithAnyMethod(),
			fcors.WithAnyRequestHeaders(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create CORS middleware: %w", err)
		}
		mux.Use(cors)
	default:
	}

	srv.Handler = api.HandlerWithOptions(handler, api.ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: mux,
	})

	return s, nil
}

func (s *server) Run(ctx context.Context) error {
	shutdownErrorChan := make(chan error)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
		defer cancel()

		shutdownErrorChan <- s.httpServer.Shutdown(ctx)
	}()

	s.logger.Info("starting server", slog.String("addr", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	s.logger.Info("stopped server", slog.String("addr", s.httpServer.Addr))

	return nil
}