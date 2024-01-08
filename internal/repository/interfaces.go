package repository

import (
	"context"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/aggregate"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/objvalue"
	"github.com/google/uuid"
)

type (
	Inbox interface {
		CreateTransaction(ctx context.Context, transaction *aggregate.Transaction) error
		UpdateTransactionStatus(ctx context.Context, transaction *aggregate.Transaction) error
		CreateTransactionHistory(ctx context.Context, transactionID uuid.UUID, history objvalue.TransactionHistory) error
		GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	}
	Outbox interface {
		GetTransaction(ctx context.Context, transactionID uuid.UUID) (*aggregate.Transaction, error)
	}

	User interface {
		Get(ctx context.Context, userID uuid.UUID) (*model.User, error)
	}

	UserCache interface {
		Get(ctx context.Context, userID uuid.UUID) (*model.User, error)
		Set(ctx context.Context, user *model.User) error
	}

	Transaction interface {
		Get(ctx context.Context, transactionID uuid.UUID) (*model.Transaction, error)
		Create(ctx context.Context, transaction *model.Transaction) error
		UpdateStatus(ctx context.Context, transaction *model.Transaction) error
		CreateHistory(ctx context.Context, transactionID uuid.UUID, history objvalue.TransactionHistory) error
	}

	TransactionCache interface {
		Get(ctx context.Context, transactionID uuid.UUID) (*model.Transaction, error)
		Set(ctx context.Context, transaction *model.Transaction) error
	}
)
