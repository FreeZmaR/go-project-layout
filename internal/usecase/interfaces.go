package usecase

import (
	"context"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/aggregate"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"
	"github.com/google/uuid"
)

type (
	Inbox interface {
		MakeTransaction(
			ctx context.Context,
			userFromID, userToID uuid.UUID,
			amount int,
			paymentMethod definition.PaymentMethod,
			currency definition.Currency,
			description string,
		) (*aggregate.Transaction, error)

		InspectTransaction(ctx context.Context, transaction *aggregate.Transaction) error
		ConfirmTransaction(ctx context.Context, transaction *aggregate.Transaction) error
	}

	Outbox interface {
		ShowTransaction(ctx context.Context, transactionID uuid.UUID) (*aggregate.Transaction, error)
	}
)
