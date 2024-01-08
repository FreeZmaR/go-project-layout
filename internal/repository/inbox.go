package repository

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/domain/aggregate"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/domain/objvalue"
	"github.com/FreeZmaR/go-project-layout/internal/lib/utils"
	"github.com/google/uuid"
	"time"
)

type inbox struct {
	trans Transaction
	user  User
}

func NewInbox(trans Transaction, userRP User) Inbox {
	return &inbox{trans: trans, user: userRP}
}

func (rp inbox) CreateTransaction(ctx context.Context, transaction *aggregate.Transaction) error {
	transModel := transaction.Model()
	transModel.CreatedAt = time.Now()

	return rp.trans.Create(ctx, transModel)
}

func (rp inbox) UpdateTransactionStatus(ctx context.Context, transaction *aggregate.Transaction) error {
	transModel := transaction.Model()
	transModel.UpdatedAt = utils.WithPtr(time.Now())

	return rp.trans.UpdateStatus(ctx, transModel)
}

func (rp inbox) CreateTransactionHistory(
	ctx context.Context,
	transactionID uuid.UUID,
	history objvalue.TransactionHistory,
) error {
	return rp.trans.CreateHistory(ctx, transactionID, history)
}

func (rp inbox) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return rp.user.Get(ctx, userID)
}
