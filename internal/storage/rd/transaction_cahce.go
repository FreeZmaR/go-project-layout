package rd

import (
	"errors"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/objvalue"
	"github.com/google/uuid"
	"time"
)

type transactionCache struct {
	TransactionID string     `json:"transaction_id"`
	PaymentMethod int        `json:"payment_method"`
	Status        int        `json:"status"`
	Amount        int        `json:"amount"`
	Currency      int        `json:"currency"`
	Description   string     `json:"description"`
	UserFromID    string     `json:"user_from_id"`
	UserToID      string     `json:"user_to_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

func (t *transactionCache) toTransaction() (model.Transaction, error) {
	id, err := uuid.Parse(t.TransactionID)
	if err != nil {
		return model.Transaction{}, errors.New("invalid transaction id")
	}

	userFromID, err := uuid.Parse(t.UserFromID)
	if err != nil {
		return model.Transaction{}, errors.New("invalid user from id")
	}

	userToID, err := uuid.Parse(t.UserToID)
	if err != nil {
		return model.Transaction{}, errors.New("invalid user to id")
	}

	if !definition.IsPaymentMethod(t.PaymentMethod) {
		return model.Transaction{}, errors.New("invalid payment method")
	}

	if !objvalue.IsTransactionStatus(t.Status) {
		return model.Transaction{}, errors.New("invalid transaction status")
	}

	if !definition.IsCurrency(t.Currency) {
		return model.Transaction{}, errors.New("invalid currency")
	}

	return model.Transaction{
		ID:            id,
		PaymentMethod: definition.PaymentMethod(t.PaymentMethod),
		Status:        objvalue.TransactionStatus(t.Status),
		Amount:        t.Amount,
		Currency:      definition.Currency(t.Currency),
		Description:   t.Description,
		UserFromID:    userFromID,
		UserToID:      userToID,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}, nil
}

func newTransactionCache(transaction *model.Transaction) transactionCache {
	return transactionCache{
		TransactionID: transaction.ID.String(),
		PaymentMethod: int(transaction.PaymentMethod),
		Status:        int(transaction.Status),
		Amount:        transaction.Amount,
		Currency:      int(transaction.Currency),
		Description:   transaction.Description,
		UserFromID:    transaction.UserFromID.String(),
		UserToID:      transaction.UserToID.String(),
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
