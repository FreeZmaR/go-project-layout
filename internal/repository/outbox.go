package repository

import (
	"context"
	"errors"
	"github.com/FreeZmaR/go-project-layout/internal/domain/aggregate"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/FreeZmaR/go-project-layout/internal/domain/objvalue"
	"github.com/FreeZmaR/go-project-layout/internal/lib/utils"
	"github.com/google/uuid"
	"log/slog"
)

type OutboxInstance struct {
	user       User
	userCache  UserCache
	trans      Transaction
	transCache TransactionCache
}

func NewOutbox(
	user User,
	uCache UserCache,
	trans Transaction,
	transCache TransactionCache,
) *OutboxInstance {
	return &OutboxInstance{
		user:       user,
		userCache:  uCache,
		trans:      trans,
		transCache: transCache,
	}
}

func (rp OutboxInstance) GetTransaction(ctx context.Context, transactionID uuid.UUID) (*aggregate.Transaction, error) {
	transModel, err := rp.getTransaction(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	userFrom, err := rp.getUser(ctx, transModel.UserFromID)
	if err != nil {
		return nil, err
	}

	userTo, err := rp.getUser(ctx, transModel.UserToID)
	if err != nil {
		return nil, err
	}

	return aggregate.NewTransaction(transModel).SetUserFrom(userFrom).SetUserTo(userTo), nil
}

func (rp OutboxInstance) getTransaction(
	ctx context.Context,
	transactionID uuid.UUID,
) (*model.Transaction, error) {
	if transaction, err := rp.transCache.Get(ctx, transactionID); nil == err {
		return transaction, nil
	}

	transaction, err := rp.trans.Get(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	if nil == transaction {
		return nil, errors.New("transaction not found")
	}

	if utils.OneOf(transaction.Status, objvalue.TransactionStatusSuccess, objvalue.TransactionStatusFail) {
		go rp.setTransactionToCache(ctx, *transaction)
	}

	return transaction, nil
}

func (rp OutboxInstance) setTransactionToCache(
	ctx context.Context,
	transaction model.Transaction,
) {
	if err := rp.transCache.Set(ctx, &transaction); err != nil {
		slog.Error("set transaction to cache error", slog.String("err", err.Error()))
	}
}

func (rp OutboxInstance) getUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	if user, err := rp.userCache.Get(ctx, userID); nil == err {
		return user, nil

	}

	user, err := rp.user.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if nil == user {
		return nil, errors.New("user not found")
	}

	return user, nil
}
