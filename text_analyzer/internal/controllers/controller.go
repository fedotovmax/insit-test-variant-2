package controllers

import (
	"log/slog"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/usecases"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	log            *slog.Logger
	analyzeUsecase *usecases.AnalyzeUsecase
}

func New(log *slog.Logger, analyzeUsecase *usecases.AnalyzeUsecase) *controller {
	return &controller{
		log:            log,
		analyzeUsecase: analyzeUsecase,
	}
}

func (c *controller) Register(router chi.Router) {
	router.Get(healthRoute, c.health)
	router.Post(analyzeRoute, c.analyze)
}
