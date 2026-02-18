package usecases

import (
	"context"
	"log/slog"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/ports"
)

type HandleEventUsecase struct {
	log     *slog.Logger
	storage ports.OperationStorage
}

func NewHandleEventUsecase(
	log *slog.Logger,
	storage ports.OperationStorage,
) *HandleEventUsecase {
	return &HandleEventUsecase{
		log:     log,
		storage: storage,
	}
}

func (u *HandleEventUsecase) Execute(ctx context.Context, newOp *domain.ToProcessOperation) error {
	//TODO:
	return nil
}
