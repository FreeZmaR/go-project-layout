package v1

import (
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/outbox/v1/handler"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(baseRouter *mux.Router, outboxUC usecase.Outbox) {
	router := baseRouter.PathPrefix("/outbox/v1").Subrouter()

	router.HandleFunc("/show-transaction/{id}", handler.ShowTransaction(outboxUC)).Methods(http.MethodGet)
}
