package ports

import (
	"context"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
)

type OperationStorage interface {
	Get(ctx context.Context, id string) (*domain.Operation, error)
	Save(ctx context.Context, id string, status domain.Status, result *domain.Result) error
}
