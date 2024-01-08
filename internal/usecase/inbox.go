package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/FreeZmaR/go-project-layout/internal/domain/aggregate"
	"github.com/FreeZmaR/go-project-layout/internal/domain/definition"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/domain/objvalue"
	"github.com/FreeZmaR/go-project-layout/internal/lib/utils"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/google/uuid"
	"log/slog"
)

type inbox struct {
	repo repository.Inbox
}

func NewInbox(inboxRepo repository.Inbox) Inbox {
	return inbox{repo: inboxRepo}
}

func (uc inbox) MakeTransaction(
	ctx context.Context,
	userFromID, userToID uuid.UUID,
	amount int,
	paymentMethod definition.PaymentMethod,
	currency definition.Currency,
	description string,
) (*aggregate.Transaction, error) {
	userFrom, err := uc.getUser(ctx, userFromID)
	if err != nil {
		return nil, fmt.Errorf("%w: user_from", err)
	}

	userTo, err := uc.getUser(ctx, userToID)
	if err != nil {
		return nil, fmt.Errorf("%w: user_to", err)
	}

	transaction := aggregate.NewTransaction(
		&model.Transaction{
			ID:            uuid.New(),
			PaymentMethod: paymentMethod,
			Status:        objvalue.TransactionStatusNew,
			Amount:        amount,
			Currency:      currency,
			Description:   description,
		},
	)
	transaction.SetUserFrom(userFrom).SetUserTo(userTo)

	if err = uc.repo.CreateTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc inbox) InspectTransaction(
	ctx context.Context,
	transaction *aggregate.Transaction,
) error {
	if transaction.Status() != objvalue.TransactionStatusNew {
		return fmt.Errorf("%w: is not new", ErrTransactionWrongStatus)
	}

	defer uc.createTransactionHistory(ctx, transaction)

	if transaction.UserFrom().Balance < transaction.Amount() {
		transaction.SetFailed(
			utils.WithPtr(definition.InsufficientFundsErrCode),
			utils.WithPtr(definition.InsufficientFundsErrCode.Description()),
		)

		return nil
	}

	transaction.SetPending()

	err := uc.repo.UpdateTransactionStatus(ctx, transaction)
	if err != nil {
		transaction.SetFailed(
			utils.WithPtr(definition.InternalErrCode),
			utils.WithPtr(definition.InternalErrCode.Description()),
		)

		return err
	}

	return nil
}

func (uc inbox) ConfirmTransaction(
	ctx context.Context,
	transaction *aggregate.Transaction,
) error {
	defer uc.createTransactionHistory(ctx, transaction)

	if transaction.Status() != objvalue.TransactionStatusPending {
		return nil
	}

	transaction.SetSuccess()

	err := uc.repo.UpdateTransactionStatus(ctx, transaction)
	if err != nil {
		transaction.SetFailed(
			utils.WithPtr(definition.InternalErrCode),
			utils.WithPtr(definition.InternalErrCode.Description()),
		)

		return err
	}

	return nil
}

func (uc inbox) getUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := uc.repo.GetUser(ctx, userID)
	if nil == err {
		return user, nil
	}

	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrUserNotFound
	}

	return nil, err
}

func (uc inbox) createTransactionHistory(
	ctx context.Context,
	transaction *aggregate.Transaction,
) {
	err := uc.repo.CreateTransactionHistory(ctx, transaction.ID(), transaction.History())
	if err != nil {
		slog.Error("failed to create transaction history", transaction.ID(), transaction.History())
	}
}
