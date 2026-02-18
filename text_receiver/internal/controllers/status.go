package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/utils"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/validation"
)

// @Summary Get operation status by id
// @Description Get operation status by id
// @Tags receiver
// @Accept json
// @Produce json
// @Param id path string true "Operation id"
// @Router /api/v1/status/{id} [get]
// @Success 200 {object} domain.Operation
// @Failure 400 {object} domain.Error
// @Failure 404 {object} domain.Error
// @Failure 500 {object} domain.Error
func (c *controller) status(w http.ResponseWriter, r *http.Request) {

	operationID := r.PathValue("id")

	err := validation.IsUUID(operationID)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, domain.Error{Message: err.Error()})
		return
	}

	getCtx, cancelGetCtx := context.WithTimeout(r.Context(), time.Second*3)
	defer cancelGetCtx()

	op, err := c.query.Get(getCtx, operationID)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, op)

}
