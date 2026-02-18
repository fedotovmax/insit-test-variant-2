package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/domain"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/domain/inputs"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/utils"
)

// @Summary      Analyze text
// @Description  Analyze text
// @Router       /api/v1/analyze [post]
// @Tags         analyzer
// @Accept       json
// @Produce      json
// @Param dto body inputs.Text true "Analyze text input"
// @Success      200  {object}  domain.AnalyzeResult
// @Failure      400  {object}  domain.Error
// @Failure      403  {object}  domain.Error
// @Failure      500  {object}  domain.Error
func (c *controller) analyze(w http.ResponseWriter, r *http.Request) {

	const op = "controllers.analyze"

	var input inputs.Text

	err := utils.DecodeJSON(r.Body, &input)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, domain.Error{Message: err.Error()})
		return
	}

	analyzeCtx, cancelAnalyzeCtx := context.WithTimeout(r.Context(), time.Second*3)
	defer cancelAnalyzeCtx()

	result := c.analyzeUsecase.Execute(analyzeCtx, input)

	utils.WriteJSON(w, http.StatusOK, result)

}
