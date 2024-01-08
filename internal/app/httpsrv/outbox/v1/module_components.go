package v1

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
)

func ProvidePostgresPoolClient(param ParamsIn) (postgres.Client, error) {
	return postgres.NewPool(param.Postgres)
}

func ProvideRedisClient(param ParamsIn) (redis.Client, error) {
	return redis.NewClient(param.Redis)
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

func InvokeInitRouter(p ParamsIn, outboxUC usecase.Outbox) {
	InitRouter(p.Router, outboxUC)
}

func InvokeFinalizer(lc fx.Lifecycle, pgClient postgres.Client, rdClient redis.Client) {
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			pgErr := pgClient.Close()
			rdErr := rdClient.Close()

			if pgErr != nil {
				return pgErr
			}

			return rdErr
		},
	})
}
