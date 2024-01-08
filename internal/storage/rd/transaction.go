package rd

import (
	"context"
	"encoding/json"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/google/uuid"
	"time"
)

const transactionKeyPrefix = "transaction:"

func GetTransaction(ctx context.Context, db redis.Connect, transactionID uuid.UUID) (*model.Transaction, error) {
	key := transactionKeyPrefix + transactionID.String()

	data, err := db.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var transCache transactionCache

	if err = json.Unmarshal([]byte(data), &transCache); err != nil {
		return nil, err
	}

	transaction, err := transCache.toTransaction()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func SetTransaction(
	ctx context.Context,
	db redis.Connect,
	transaction *model.Transaction,
	expiration time.Duration,
) error {
	transCache := newTransactionCache(transaction)

	data, err := json.Marshal(transCache)
	if err != nil {
		return err
	}

	key := transactionKeyPrefix + transaction.ID.String()

	return db.Set(ctx, key, string(data), expiration)
}
