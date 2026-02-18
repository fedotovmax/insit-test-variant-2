package usecases

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/ports"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/queries"
)

type SaveUsecase struct {
	log       *slog.Logger
	query     queries.Operation
	storage   ports.OperationStorage
	publisher ports.EventPublisher
}

func NewSaveUsecase(log *slog.Logger,
	query queries.Operation,
	storage ports.OperationStorage,
	publisher ports.EventPublisher,
) *SaveUsecase {
	return &SaveUsecase{
		storage:   storage,
		log:       log,
		query:     query,
		publisher: publisher,
	}
}

func (u *SaveUsecase) Execute(
	ctx context.Context,
	id string,
	text string,
	status domain.Status,
	result *domain.Result,
) error {

	const op = "usecases.save"

	err := u.storage.Save(ctx, id, status, result)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = u.publisher.Send(&domain.ToProcessOperation{
		ID:   id,
		Text: text,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
