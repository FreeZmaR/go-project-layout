package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/FreeZmaR/go-project-layout/internal/storage/rd"
	"github.com/google/uuid"
)

const (
	transactionExpirationTime = 60 * 60 * 24 * 7
)

type transactionCache struct {
	db redis.Connect
}

func NewTransactionCache(db redis.Connect) TransactionCache {
	return &transactionCache{db: db}
}

func (rp transactionCache) Get(ctx context.Context, transactionID uuid.UUID) (*model.Transaction, error) {
	return rd.GetTransaction(ctx, rp.db, transactionID)
}

func (rp transactionCache) Set(ctx context.Context, transaction *model.Transaction) error {
	return rd.SetTransaction(ctx, rp.db, transaction, transactionExpirationTime)
}
