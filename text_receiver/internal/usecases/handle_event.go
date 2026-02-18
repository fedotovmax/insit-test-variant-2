package usecases

import (
	"context"
	"log/slog"
	"time"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/clients/http/analyzer"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/ports"
)

type HandleEventUsecase struct {
	log            *slog.Logger
	analyzerClient *analyzer.APIClient
	storage        ports.OperationStorage
}

func NewHandleEventUsecase(
	log *slog.Logger,
	analyzerClient *analyzer.APIClient,
	storage ports.OperationStorage,
) *HandleEventUsecase {
	return &HandleEventUsecase{
		log:            log,
		analyzerClient: analyzerClient,
		storage:        storage,
	}
}

func (u *HandleEventUsecase) Execute(ctx context.Context, newOp *domain.ToProcessOperation) error {

	analyzeCtx, cancelAnalyzeCtx := context.WithTimeout(ctx, time.Second*3)
	defer cancelAnalyzeCtx()

	analyzeResult, _, err := u.analyzerClient.AnalyzerAPI.
		ApiV1AnalyzePost(analyzeCtx).
		Dto(analyzer.InputsText{Data: newOp.Text}).
		Execute()

	if err != nil {
		return err
	}

	saveCtx, cancelSaveCtx := context.WithTimeout(ctx, time.Second*3)
	defer cancelSaveCtx()

	res := &domain.Result{
		WordCount:         int(analyzeResult.WordCount),
		SymbolsCount:      int(analyzeResult.SymbolsCount),
		SentencesCount:    int(analyzeResult.SentencesCount),
		AverageWordLength: float64(analyzeResult.AverageWordLength),
	}

	err = u.storage.Save(saveCtx, newOp.ID, domain.StatusCompleted, res)

	if err != nil {
		return err
	}

	return nil
}
