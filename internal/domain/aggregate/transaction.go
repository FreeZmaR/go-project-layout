package aggregate

import (
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/model"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/objvalue"
	"github.com/google/uuid"
)

type Transaction struct {
	transaction *model.Transaction
	userFrom    *model.User
	userTo      *model.User
	history     objvalue.TransactionHistory
}

func NewTransaction(transaction *model.Transaction) *Transaction {
	return &Transaction{
		transaction: transaction,
	}
}

func (t *Transaction) SetUserFrom(user *model.User) *Transaction {
	t.userFrom = user
	t.transaction.UserFromID = user.ID

	return t
}

func (t *Transaction) SetUserTo(user *model.User) *Transaction {
	t.userTo = user
	t.transaction.UserToID = user.ID

	return t
}

func (t *Transaction) SetFailed(errCode *definition.ErrorCode, description *string) *Transaction {
	t.setStatus(objvalue.TransactionStatusFail)
	t.history.ErrorCode = errCode
	t.history.ErrorDescription = description

	return t
}

func (t *Transaction) SetSuccess() *Transaction {
	t.setStatus(objvalue.TransactionStatusSuccess)

	return t
}

func (t *Transaction) SetPending() *Transaction {
	t.setStatus(objvalue.TransactionStatusPending)

	return t
}

func (t *Transaction) Model() *model.Transaction {
	return t.transaction
}

func (t *Transaction) ID() uuid.UUID {
	return t.transaction.ID
}

func (t *Transaction) Status() objvalue.TransactionStatus {
	return t.transaction.Status
}

func (t *Transaction) Amount() int {
	return t.transaction.Amount
}

func (t *Transaction) Currency() definition.Currency {
	return t.transaction.Currency
}

func (t *Transaction) PaymentMethod() definition.PaymentMethod {
	return t.transaction.PaymentMethod
}

func (t *Transaction) UserFrom() *model.User {
	return t.userFrom
}

func (t *Transaction) UserTo() *model.User {
	return t.userTo
}

func (t *Transaction) History() objvalue.TransactionHistory {
	return t.history
}

func (t *Transaction) ErrorCode() *definition.ErrorCode {
	return t.history.ErrorCode
}

func (t *Transaction) ErrorDescription() *string {
	return t.history.ErrorDescription
}

func (t *Transaction) setStatus(status objvalue.TransactionStatus) {
	t.history.OldStatus = t.transaction.Status
	t.transaction.Status = status
	t.history.NewStatus = t.transaction.Status
}
