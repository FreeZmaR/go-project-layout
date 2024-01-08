package usecase

import "errors"

var (
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrTransactionWrongStatus = errors.New("transaction has wrong status")
	ErrUserNotFound           = errors.New("user not found")
)
