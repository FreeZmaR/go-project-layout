package v1

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/lib/fxutils"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
)

func ProvidePostgresPoolClient(param ParamsIn, finalizer *fxutils.Finalizer) (postgres.Client, error) {
	pool, err := postgres.NewPool(param.Postgres)
	if err != nil {
		return nil, err
	}

	finalizer.Append(pool)

	return pool, nil
}

func ProvideRedisClient(param ParamsIn, finalizer *fxutils.Finalizer) (redis.Client, error) {
	client, err := redis.NewClient(param.Redis)
	if err != nil {
		return nil, err
	}

	finalizer.Append(client)

	return client, nil
}

func ProvideOutboxUseCase(repo repository.Outbox) usecase.Outbox {
	return usecase.NewOutbox(repo)
}

func ProvideOutboxRepository(
	userRP repository.User,
	userCacheRP repository.UserCache,
	transRP repository.Transaction,
	transCacheRP repository.TransactionCache,
) repository.Outbox {
	return repository.NewOutbox(userRP, userCacheRP, transRP, transCacheRP)
}

func ProvideTransactionRepository(pgClient postgres.Client) repository.Transaction {
	return repository.NewTransaction(pgClient)
}

func ProvideUserRepository(pgClient postgres.Client) repository.User {
	return repository.NewUser(pgClient)
}

func ProvideUserCacheRepository(rdClient redis.Client) repository.UserCache {
	return repository.NewUserCache(rdClient)
}

func ProvideTransactionCacheRepository(rdClient redis.Client) repository.TransactionCache {
	return repository.NewTransactionCache(rdClient)
}

func ProvideFinaliser() *fxutils.Finalizer {
	return fxutils.NewFinalizer()
}

func InvokeInitRouter(p ParamsIn, outboxUC usecase.Outbox) {
	InitRouter(p.Router, outboxUC)
}

func InvokeFinalizer(lc fx.Lifecycle, finalizer *fxutils.Finalizer) {
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			finalizer.Close()

			return nil
		},
	})
}
