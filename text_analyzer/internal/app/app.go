package app

import (
	"context"
	"fmt"
	"log/slog"

	_ "github.com/fedotovmax74/insit-test-variant-2/text_analyzer/docs"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/adapters/server/http"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/config"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/controllers"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/usecases"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/pkg/logger"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type App struct {
	appConfig  *config.AppConfig
	log        *slog.Logger
	httpServer *http.Server
}

func New(appConfig *config.AppConfig, log *slog.Logger) (*App, error) {

	const op = "app.New"

	router := chi.NewRouter()

	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	analyzeUsecase := usecases.NewAnalyzeUsecase(log)

	controller := controllers.New(log, analyzeUsecase)

	controller.Register(router)

	httpServer := http.New(appConfig.HTTPServer, router)

	return &App{
		appConfig:  appConfig,
		log:        log,
		httpServer: httpServer,
	}, nil
}

func (a *App) Start() <-chan error {
	const op = "app.Start"

	log := a.log.With(slog.String("op", op))

	errChan := make(chan error, 1)

	go func() {
		log.Info(
			"Starting HTTP server...",
			slog.String("addr", fmt.Sprintf("http://localhost:%d", a.appConfig.HTTPServer.Port)),
		)
		if err := a.httpServer.Start(); err != nil {
			errChan <- fmt.Errorf("%s: %w", op, err)
		}
	}()

	return errChan
}

func (a *App) Stop(ctx context.Context) {
	const op = "app.Start"

	log := a.log.With(slog.String("op", op))

	if err := a.httpServer.Stop(ctx); err != nil {
		log.Error("Error when shutdown HTTP server", logger.Err(err))
	} else {
		log.Info("HTTP server stopped successfully!")
	}
}
