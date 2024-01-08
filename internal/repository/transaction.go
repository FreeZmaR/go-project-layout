package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/domain/objvalue"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/storage/pg"
	"github.com/google/uuid"
	"time"
)

type transactionRP struct {
	db postgres.Connect
}

func NewTransaction(db postgres.Connect) Transaction {
	return &transactionRP{db: db}
}

func (rp transactionRP) Get(ctx context.Context, transactionID uuid.UUID) (*model.Transaction, error) {
	return pg.GetTransaction(ctx, rp.db, transactionID)
}

func (rp transactionRP) Create(ctx context.Context, transaction *model.Transaction) error {
	return pg.CreateTransaction(ctx, rp.db, transaction)
}

func (rp transactionRP) UpdateStatus(ctx context.Context, transaction *model.Transaction) error {
	return pg.UpdateTransactionStatus(ctx, rp.db, transaction)
}

func (rp transactionRP) CreateHistory(
	ctx context.Context,
	transactionID uuid.UUID,
	history objvalue.TransactionHistory,
) error {
	return pg.CreateTransactionHistory(ctx, rp.db, transactionID, history, time.Now())
}
