package operation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/db/redis"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
)

func (r *redisDb) Save(ctx context.Context, id string, status domain.Status, res *domain.Result) error {

	const op = "adapters.db.redis.operation.Save"

	operation := domain.Operation{
		ID:     id,
		Status: status,
		Result: res,
	}

	bytes, err := json.Marshal(operation)

	if err != nil {
		return fmt.Errorf("%s: %w: %v", op, adapters.ErrInternal, err)
	}

	_, err = r.rdb.Set(ctx, redis.Operation(id), bytes, 0).Result()

	if err != nil {
		return fmt.Errorf("%s: %w: %v", op, adapters.ErrInternal, err)
	}

	return nil
}
