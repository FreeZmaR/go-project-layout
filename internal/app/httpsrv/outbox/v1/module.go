package v1

import (
	"github.com/FreeZmaR/go-project-layout/config/types"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

const moduleName = "outbox-v1"

type ParamsIn struct {
	fx.In

	Postgres *types.Postgres
	Redis    *types.Redis
	Router   *mux.Router
}

func NewModule() fx.Option {
	return fx.Module(
		moduleName,
		fx.Provide(
			ProvideFinaliser,
			ProvidePostgresPoolClient,
			ProvideRedisClient,
			ProvideOutboxUseCase,
			ProvideOutboxRepository,
			ProvideUserRepository,
			ProvideUserCacheRepository,
			ProvideTransactionRepository,
			ProvideTransactionCacheRepository,
		),
		fx.Invoke(
			InvokeFinalizer,
			InvokeInitRouter,
		),
	)
}
