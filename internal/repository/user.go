package repository

import (
	"context"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/lib/postgres"
	"github.com/FreeZmaR/go-service-structure/template/internal/storage/pg"
	"github.com/google/uuid"
)

type userRP struct {
	db postgres.Connect
}

func NewUser(db postgres.Connect) User {
	return &userRP{db: db}
}

func (rp userRP) Get(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return pg.GetUser(ctx, rp.db, userID)
}
