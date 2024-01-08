package pg

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/google/uuid"
)

func GetUser(ctx context.Context, db postgres.Connect, userID uuid.UUID) (*model.User, error) {
	sql := `
		SELECT
			id,
			name,
			age,
			balance,
			available_payment_methods
		FROM users
		WHERE id = $1
	`

	row := db.QueryRow(ctx, sql, userID)

	user := model.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.Balance,
		&user.AvailablePaymentMethods,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
