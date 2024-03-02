package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/lib/redis"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"go.uber.org/fx"
)

func ProvideRepositories() fx.Option {
	return fx.Provide(
		provideOutboxRP,
		provideTransactionRP,
		provideUserRP,
		provideUserCacheRP,
		provideTransactionCacheRP,
	)
}

func provideOutboxRP(
	userRP repository.User,
	userCacheRP repository.UserCache,
	transRP repository.Transaction,
	transCacheRP repository.TransactionCache,
) repository.Outbox {
	return repository.NewOutbox(userRP, userCacheRP, transRP, transCacheRP)
}

func provideTransactionRP(pgClient postgres.Client) repository.Transaction {
	return repository.NewTransaction(pgClient)
}

func provideUserRP(pgClient postgres.Client) repository.User {
	return repository.NewUser(pgClient)
}

func provideUserCacheRP(rdClient redis.Client) repository.UserCache {
	return repository.NewUserCache(rdClient)
}

func provideTransactionCacheRP(rdClient redis.Client) repository.TransactionCache {
	return repository.NewTransactionCache(rdClient)
}
