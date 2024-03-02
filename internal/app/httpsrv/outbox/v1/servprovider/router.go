package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/outbox/v1/handler"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
	"net/http"
)

func ProvideRouter() fx.Option {
	return fx.Provide(provideInitRouter)
}

func provideInitRouter(p ParamsIn, outboxUC usecase.Outbox) {
	router := p.Router.PathPrefix("/outbox/v1").Subrouter()

	router.HandleFunc("/show-transaction/{id}", handler.ShowTransaction(outboxUC)).Methods(http.MethodGet)
}
