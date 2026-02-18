package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain/inputs"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/utils"
	"github.com/google/uuid"
)

// @Summary      Receive new text for analyze
// @Description  Receive new text for analyze
// @Router       /api/v1/text [post]
// @Tags         receiver
// @Accept       json
// @Produce      json
// @Param dto body inputs.NewOperation true "Receive new text input"
// @Success      200  {object}  domain.SaveResponse
// @Failure      400  {object}  domain.Error
// @Failure      403  {object}  domain.Error
// @Failure      500  {object}  domain.Error
func (c *controller) text(w http.ResponseWriter, r *http.Request) {
	//todo: описать сваггер нормально
	const op = "controllers.text"

	var input inputs.NewOperation

	err := utils.DecodeJSON(r.Body, &input)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, domain.Error{Message: err.Error()})
		return
	}

	operationID := uuid.New().String()

	prepareOperationSaveCtx, cancelPrepareOperationSaveCtx := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelPrepareOperationSaveCtx()

	err = c.saveUsecase.Execute(prepareOperationSaveCtx, operationID, input.Text, domain.StatusInProcess, nil)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, domain.SaveResponse{ID: operationID})

}
