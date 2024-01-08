package model

import (
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/objvalue"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uuid.UUID
	PaymentMethod definition.PaymentMethod
	Status        objvalue.TransactionStatus
	Amount        int
	Currency      definition.Currency
	Description   string
	UserFromID    uuid.UUID
	UserToID      uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
