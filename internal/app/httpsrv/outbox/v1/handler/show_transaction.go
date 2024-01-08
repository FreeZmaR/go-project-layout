package handler

import (
	"errors"
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/outbox/v1/responder"
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/outbox/v1/validation"
	"github.com/FreeZmaR/go-service-structure/template/internal/usecase"
	"log/slog"
	"net/http"
)

func ShowTransaction(outboxUC usecase.Outbox) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("begin show transaction")
		rsp := responder.NewShowTransaction(w)

		input, err := validation.ValidateShowTransactionRequest(r)
		if err != nil {
			slog.Info("validation error", slog.String("err", err.Error()))
			rsp.ValidationError(err)

			return
		}

		transaction, err := outboxUC.ShowTransaction(r.Context(), input.TransactionID)
		if err != nil {
			slog.Info("error show transaction", slog.String("err", err.Error()))

			if errors.Is(err, usecase.ErrTransactionNotFound) {
				rsp.NotFound()

				return
			}

			rsp.InternalError()

			return
		}

		slog.Info("found transaction", slog.Any("transaction", transaction))
		rsp.SendTransaction(transaction)
	}
}
