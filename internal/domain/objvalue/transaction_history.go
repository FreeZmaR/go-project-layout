package objvalue

import "github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"

type TransactionHistory struct {
	OldStatus        TransactionStatus
	NewStatus        TransactionStatus
	ErrorCode        *definition.ErrorCode
	ErrorDescription *string
}
