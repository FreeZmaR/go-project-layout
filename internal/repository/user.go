package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/storage/pg"
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
