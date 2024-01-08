package v1

import (
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

const moduleName = "inbox-v1"

type ParamsIn struct {
	fx.In

	Postgres *types.Postgres
	Router   *mux.Router
}

func NewModule() fx.Option {
	return fx.Module(
		moduleName,
		fx.Provide(
			ProvidePostgresPoolClient,
			ProvideTransactionRepository,
			ProvideUserRepository,
			ProvideInboxRepository,
			ProvideInboxUseCase,
		),
		fx.Invoke(
			InvokeInitRouter,
			InvokeFinalizer,
		),
	)
}
