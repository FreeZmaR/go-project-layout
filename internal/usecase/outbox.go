package usecase

import (
	"context"
	"errors"
	"github.com/FreeZmaR/go-project-layout/internal/domain/aggregate"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/google/uuid"
)

type outbox struct {
	repo repository.Outbox
}

func NewOutbox(repo repository.Outbox) Outbox {
	return outbox{repo: repo}
}

func (uc outbox) ShowTransaction(ctx context.Context, transactionID uuid.UUID) (*aggregate.Transaction, error) {
	transaction, err := uc.repo.GetTransaction(ctx, transactionID)
	if nil == err {
		return transaction, nil
	}

	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrTransactionNotFound
	}

	return nil, err
}
