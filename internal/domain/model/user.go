package model

import (
	"github.com/FreeZmaR/go-project-layout/internal/domain/definition"
	"github.com/google/uuid"
)

type User struct {
	ID                      uuid.UUID
	Name                    string
	Age                     int
	Balance                 int
	AvailablePaymentMethods []definition.PaymentMethod
}
