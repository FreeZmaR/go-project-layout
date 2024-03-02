package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/inbox/v1/handler"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
	"net/http"
)

func ProvideRouter() fx.Option {
	return fx.Invoke(
		provideInitRouter,
	)
}

func provideInitRouter(p ParamsIn, inboxUC usecase.Inbox) {
	router := p.Router.PathPrefix("/inbox/v1").Subrouter()

	router.HandleFunc("/new-transaction", handler.NewTransaction(inboxUC)).Methods(http.MethodPost)
}
