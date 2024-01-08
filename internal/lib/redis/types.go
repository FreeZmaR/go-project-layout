package redis

import (
	"context"
	"time"
)

type Connect interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
