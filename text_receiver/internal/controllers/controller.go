package controllers

import (
	"log/slog"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/queries"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/usecases"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	log         *slog.Logger
	saveUsecase *usecases.SaveUsecase
	query       queries.Operation
}

func New(
	log *slog.Logger,
	saveUsecase *usecases.SaveUsecase,
	query queries.Operation,
) *controller {
	return &controller{
		log:         log,
		saveUsecase: saveUsecase,
		query:       query,
	}
}

func (c *controller) Register(router chi.Router) {
	router.Get(healthRoute, c.health)
	router.Post(textRoute, c.text)
	router.Get(statusIDRoute, c.status)
}
