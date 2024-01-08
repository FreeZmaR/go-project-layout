package validation

import (
	"encoding/json"
	"errors"
	"github.com/FreeZmaR/go-project-layout/internal/domain/definition"
	"github.com/FreeZmaR/go-project-layout/internal/lib/utils"
	"github.com/google/uuid"
	"net/http"
)

type NewTransactionValidated struct {
	Amount        int
	PaymentMethod definition.PaymentMethod
	Currency      definition.Currency
	Description   string
	UserFrom      uuid.UUID
	UserTo        uuid.UUID
}

type newTransactionInput struct {
	Amount        int    `json:"amount"`
	PaymentMethod int    `json:"payment_method"`
	Currency      int    `json:"currency"`
	Description   string `json:"description"`
	UserFrom      string `json:"user_from"`
	UserTo        string `json:"user_to"`
}

func ValidateNewTransactionRequest(r *http.Request) (*NewTransactionValidated, error) {
	var input newTransactionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, errors.New("invalid request body")
	}

	if input.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	paymentMethod := definition.PaymentMethod(input.PaymentMethod)

	if utils.OneOf(
		paymentMethod,
		definition.CardPaymentMethod,
		definition.PayPalPaymentMethod,
		definition.WalletPaymentMethod,
	) {
		return nil, errors.New("invalid payment_method")
	}

	currency := definition.Currency(input.Currency)

	if utils.OneOf(currency, definition.RUB, definition.USD, definition.EUR) {
		return nil, errors.New("invalid currency")
	}

	userFrom, err := uuid.Parse(input.UserFrom)
	if err != nil {
		return nil, errors.New("invalid user_from")
	}

	userTo, err := uuid.Parse(input.UserTo)
	if err != nil {
		return nil, errors.New("invalid user_to")
	}

	return &NewTransactionValidated{
		Amount:        input.Amount,
		PaymentMethod: paymentMethod,
		Currency:      currency,
		Description:   input.Description,
		UserFrom:      userFrom,
		UserTo:        userTo,
	}, nil
}
