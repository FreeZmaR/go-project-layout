package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/FreeZmaR/go-project-layout/internal/storage/rd"
	"github.com/google/uuid"
)

const userExpirationTime = 60 * 60 * 24 * 7

type userCache struct {
	db redis.Connect
}

func NewUserCache(db redis.Connect) UserCache {
	return &userCache{db: db}
}

func (rp userCache) Get(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return rd.GetUser(ctx, rp.db, userID)
}

func (rp userCache) Set(ctx context.Context, user *model.User) error {
	return rd.SetUser(ctx, rp.db, *user, userExpirationTime)
}
