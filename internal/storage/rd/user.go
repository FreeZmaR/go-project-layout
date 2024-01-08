package rd

import (
	"context"
	"encoding/json"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/lib/redis"
	"github.com/google/uuid"
	"time"
)

const userKeyPrefix = "user:"

func GetUser(ctx context.Context, db redis.Connect, userID uuid.UUID) (*model.User, error) {
	key := userKeyPrefix + userID.String()

	data, err := db.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var uCache userCache

	if err = json.Unmarshal([]byte(data), &uCache); err != nil {
		return nil, err
	}

	user, err := uCache.ToUser()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SetUser(
	ctx context.Context,
	db redis.Connect,
	user model.User,
	expiration time.Duration,
) error {
	uCache := newUserCache(user)

	data, err := json.Marshal(uCache)
	if err != nil {
		return err
	}

	key := userKeyPrefix + user.ID.String()

	return db.Set(ctx, key, string(data), expiration)
}
