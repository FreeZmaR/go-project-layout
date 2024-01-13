package v1

import (
	"context"
	"github.com/FreeZmaR/go-project-layout/internal/lib/fxutils"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
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

func ProvideFinalizer() *fxutils.Finalizer {
	return fxutils.NewFinalizer()
}

func InvokeInitRouter(p ParamsIn, inboxUC usecase.Inbox) {
	InitRouter(p.Router, inboxUC)
}

func InvokeFinalizer(lc fx.Lifecycle, finalizer *fxutils.Finalizer) {
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			finalizer.Close()

			return nil
		},
	})
}
