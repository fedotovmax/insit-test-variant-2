package controllers

import (
	"net/http"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/utils"
)

// @Summary      health check
// @Description  health check
// @Router       /api/v1/health [get]
// @Tags         receiver
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Error
// @Failure      500  {object}  domain.Error
func (c *controller) health(w http.ResponseWriter, r *http.Request) {

	const op = "controllers.health"

	utils.WriteJSON(w, http.StatusOK, domain.Error{Message: "OK"})

}
