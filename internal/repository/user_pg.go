package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/storage/pg"
	"github.com/google/uuid"
)

type UserPG struct {
	db postgres.Connect
}

func NewUser(db postgres.Connect) *UserPG {
	return &UserPG{db: db}
}

func (rp UserPG) Get(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return pg.GetUser(ctx, rp.db, userID)
}
