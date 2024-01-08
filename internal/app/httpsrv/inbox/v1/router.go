package v1

import (
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/inbox/v1/handler"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(baseRouter *mux.Router, inboxUC usecase.Inbox) {
	router := baseRouter.PathPrefix("/inbox/v1").Subrouter()

	router.HandleFunc("/new-transaction", handler.NewTransaction(inboxUC)).Methods(http.MethodPost)
}
