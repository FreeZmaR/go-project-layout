package pg

import (
	"context"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/lib/postgres"
	"github.com/google/uuid"
)

func CreateTransaction(
	ctx context.Context,
	conn postgres.Connect,
	transaction *model.Transaction,
) error {
	sql := `
INSERT INTO transactions (
	id,
	payment_method,
	amount,
	currency,
	status,
	description,
	user_from_id,
	user_to_id,
	created_at,
 	updated_at
) VALUES ($1, $2, $3, $4, $5, $6)`

	return conn.Exec(
		ctx,
		sql,
		transaction.ID,
		transaction.PaymentMethod,
		transaction.Amount,
		transaction.Currency,
		transaction.Status,
		transaction.Description,
		transaction.UserFromID,
		transaction.UserToID,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
}

func UpdateTransactionStatus(
	ctx context.Context,
	conn postgres.Connect,
	transaction *model.Transaction,
) error {
	sql := `
UPDATE transactions
SET status = $1, updated_at = $2
WHERE id = $3`

	return conn.Exec(
		ctx,
		sql,
		transaction.Status,
		transaction.UpdatedAt,
		transaction.ID,
	)
}

func GetTransaction(
	ctx context.Context,
	conn postgres.Connect,
	transactionID uuid.UUID,
) (*model.Transaction, error) {
	sql := `
SELECT 
	payment_method,
	amount,
	currency,
	status,
	description,
	user_from_id,
	user_to_id,
	created_at,
	updated_at
FROM transactions
WHERE id = $1`

	transaction := model.Transaction{ID: transactionID}

	err := conn.QueryRow(ctx, sql, transactionID).
		Scan(
			&transaction.PaymentMethod,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Status,
			&transaction.Description,
			&transaction.UserFromID,
			&transaction.UserToID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
