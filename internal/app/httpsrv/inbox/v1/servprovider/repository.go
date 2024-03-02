package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"go.uber.org/fx"
)

func ProvideRepositories() fx.Option {
	return fx.Provide(
		provideInboxRP,
		provideTransactionRP,
		provideUserRP,
	)
}

func provideInboxRP(transRP repository.Transaction, userRP repository.User) repository.Inbox {
	return repository.NewInbox(transRP, userRP)
}

func provideTransactionRP(pgClient postgres.Client) repository.Transaction {
	return repository.NewTransaction(pgClient)
}

func provideUserRP(pgClient postgres.Client) repository.User {
	return repository.NewUser(pgClient)
}
