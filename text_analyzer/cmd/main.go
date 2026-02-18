package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/app"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/config"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/pkg/logger"
)

func setupLooger(env config.AppEnv) *slog.Logger {
	if env == config.Development {
		return logger.NewHandler(slog.LevelDebug)
	}
	return logger.NewHandler(slog.LevelWarn)
}

// @title Swagger Documentation for text_analyzer app
// @version 1.0
// @description Swagger Documentation for text_analyzer app (тестовое задание для компании "Инсит")
// @contact.name Fedotv Maxim (developer)
// @contact.email f3d0t0tvmax@yandex.ru
func main() {
	appConfig, err := config.New()

	if err != nil {
		logger.GetFallback().Error(err.Error())
		os.Exit(1)
	}

	log := setupLooger(appConfig.Env)

	app, err := app.New(appConfig, log)

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	notifyCtx, cancelNotifyCtx := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancelNotifyCtx()

	startErrChan := app.Start()

	select {
	case err := <-startErrChan:
		log.Error("Error recivied when starting application", logger.Err(err))
		cancelNotifyCtx()
	case <-notifyCtx.Done():
		log.Info("OS signal recevied")
	}

	log.Info("Starting to shutdown all resources")

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdownCtx()

	app.Stop(shutdownCtx)

}
