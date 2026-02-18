package controllers

import (
	"net/http"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/domain"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/utils"
)

// @Summary      health check
// @Description  health check
// @Router       /api/v1/health [get]
// @Tags         analyzer
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Error
// @Failure      500  {object}  domain.Error
func (c *controller) health(w http.ResponseWriter, r *http.Request) {

	const op = "controllers.health"

	utils.WriteJSON(w, http.StatusOK, domain.Error{Message: "OK"})

}
