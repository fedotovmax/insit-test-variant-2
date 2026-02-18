package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/clients/http/analyzer"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/db/redis"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/db/redis/operation"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/server/http"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/config"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/controllers"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/events"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/queries"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/usecases"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/pkg/logger"
	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type App struct {
	appConfig    *config.AppConfig
	log          *slog.Logger
	httpServer   *http.Server
	redis        *redis.RedisDb
	eventManager *events.Manager
}

func New(appConfig *config.AppConfig, log *slog.Logger) (*App, error) {

	const op = "app.New"

	l := log.With(slog.String("op", op))

	rdb, err := redis.New(*appConfig.Redis, log)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	l.Info("redis client successfully connected")

	redisClient := rdb.GetClient()

	analyzerClientConfig := analyzer.NewConfiguration()

	analyzerClientConfig.Servers = []analyzer.ServerConfiguration{
		{URL: appConfig.AnalyzerClientURL},
	}

	analyzerClient := analyzer.NewAPIClient(analyzerClientConfig)
	_ = analyzerClient

	operationRedis := operation.New(log, redisClient)

	operationQuery := queries.NewOperation(operationRedis)

	handleEventUsecase := usecases.NewHandleEventUsecase(
		log,
		analyzerClient,
		operationRedis,
	)

	eventManager := events.New(100, 5, handleEventUsecase.Execute, log)

	saveUsecase := usecases.NewSaveUsecase(log, operationQuery, operationRedis, eventManager)

	router := chi.NewRouter()

	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	controller := controllers.New(log, saveUsecase, operationQuery)

	controller.Register(router)

	httpServer := http.New(appConfig.HTTPServer, router)

	return &App{
		appConfig:    appConfig,
		log:          log,
		httpServer:   httpServer,
		eventManager: eventManager,
		redis:        rdb,
	}, nil
}

func (a *App) Start() <-chan error {
	const op = "app.Start"

	log := a.log.With(slog.String("op", op))

	a.eventManager.Start()

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

	if err := a.eventManager.Stop(ctx); err != nil {
		log.Error("Error when stop event manager", logger.Err(err))
	} else {
		log.Info("Event manager stopped successfully!")
	}

	if err := a.redis.Stop(ctx); err != nil {
		log.Error("Error when stop redis client", logger.Err(err))
	} else {
		log.Info("Redis client stopped successfully!")
	}
}
