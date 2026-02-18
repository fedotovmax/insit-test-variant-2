package queries

import (
	"context"
	"errors"
	"fmt"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain/errs"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/ports"
)

type Operation interface {
	Get(ctx context.Context, id string) (*domain.Operation, error)
}

type operation struct {
	storage ports.OperationStorage
}

func NewOperation(storage ports.OperationStorage) Operation {
	return &operation{storage: storage}
}

func (q *operation) Get(ctx context.Context, id string) (*domain.Operation, error) {
	status, err := q.storage.Get(ctx, id)

	if err != nil {
		if errors.Is(err, adapters.ErrNotFound) {
			return nil, fmt.Errorf("%w: %v", errs.ErrOperationNotFound, err)
		}
		return nil, fmt.Errorf("%w", err)
	}

	return status, nil
}
