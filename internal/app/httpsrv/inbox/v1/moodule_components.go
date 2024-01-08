package v1

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
)

func ProvidePostgresPoolClient(param ParamsIn) (postgres.Client, error) {
	return postgres.NewPool(param.Postgres)
}

func ProvideInboxUseCase(inboxRP repository.Inbox) usecase.Inbox {
	return usecase.NewInbox(inboxRP)
}

func ProvideInboxRepository(transRP repository.Transaction, userRP repository.User) repository.Inbox {
	return repository.NewInbox(transRP, userRP)
}

func ProvideTransactionRepository(pgClient postgres.Client) repository.Transaction {
	return repository.NewTransaction(pgClient)
}

func ProvideUserRepository(pgClient postgres.Client) repository.User {
	return repository.NewUser(pgClient)
}

func InvokeInitRouter(p ParamsIn, inboxUC usecase.Inbox) {
	InitRouter(p.Router, inboxUC)
}

func InvokeFinalizer(lc fx.Lifecycle, client postgres.Client) {
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return client.Close()
		},
	})
}
