package validation

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type GetTransactionValidated struct {
	TransactionID uuid.UUID
}

func ValidateShowTransactionRequest(r *http.Request) (*GetTransactionValidated, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, errors.New("transaction id not provided")
	}

	transactionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &GetTransactionValidated{
		TransactionID: transactionID,
	}, nil
}
