package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/config/types"
	"github.com/FreeZmaR/go-project-layout/internal/lib/fxutils"
	"github.com/FreeZmaR/go-project-layout/internal/lib/postgres"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type ParamsIn struct {
	fx.In

	Postgres *types.Postgres
	Router   *mux.Router
}

func ProvideServices() fx.Option {
	return fx.Provide(
		providePostgresClient,
	)
}

func providePostgresClient(param ParamsIn, finalizer *fxutils.Finalizer) (postgres.Client, error) {
	pool, err := postgres.NewPool(param.Postgres)
	if err != nil {
		return nil, err
	}

	finalizer.Append(pool)

	return pool, nil
}
