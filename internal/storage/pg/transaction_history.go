package pg

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/objvalue"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/google/uuid"
	"time"
)

func CreateTransactionHistory(
	ctx context.Context,
	db postgres.Connect,
	transactionID uuid.UUID,
	history objvalue.TransactionHistory,
	date time.Time,
) error {
	sql := `
INSERT INTO transaction_histories (
	transaction_id,
	old_status,
	new_status,
	error_code,
	error_description,
	created_at
) VALUES ($1, $2, $3, $4, $5, $6)`

	return db.Exec(
		ctx,
		sql,
		transactionID,
		history.OldStatus,
		history.NewStatus,
		history.ErrorCode,
		history.ErrorDescription,
		date,
	)
}
