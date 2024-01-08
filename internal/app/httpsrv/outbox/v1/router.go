package v1

import (
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/outbox/v1/handler"
	"github.com/FreeZmaR/go-service-structure/template/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(baseRouter *mux.Router, outboxUC usecase.Outbox) {
	router := baseRouter.PathPrefix("/outbox/v1").Subrouter()

	router.HandleFunc("/show-transaction/{id}", handler.ShowTransaction(outboxUC)).Methods(http.MethodGet)
}
