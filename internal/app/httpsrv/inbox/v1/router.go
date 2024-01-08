package v1

import (
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/inbox/v1/handler"
	"github.com/FreeZmaR/go-service-structure/template/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(baseRouter *mux.Router, inboxUC usecase.Inbox) {
	router := baseRouter.PathPrefix("/inbox/v1").Subrouter()

	router.HandleFunc("/new-transaction", handler.NewTransaction(inboxUC)).Methods(http.MethodPost)
}
