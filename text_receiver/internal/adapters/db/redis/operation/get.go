package operation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/adapters/db/redis"
	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"

	goredis "github.com/redis/go-redis/v9"
)

func (r *redisDb) Get(ctx context.Context, id string) (*domain.Operation, error) {

	const op = "adapters.db.redis.operation.Get"

	bytes, err := r.rdb.Get(ctx, redis.Operation(id)).Bytes()

	if err != nil {
		if err == goredis.Nil {
			return nil, fmt.Errorf("%s: %w: %v", op, adapters.ErrNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w: %v", op, adapters.ErrInternal, err)
	}

	var operation domain.Operation

	err = json.Unmarshal(bytes, &operation)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &operation, nil
}
