package handler

import (
	"errors"
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/inbox/v1/responder"
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/inbox/v1/validation"
	"github.com/FreeZmaR/go-service-structure/template/internal/usecase"
	"log/slog"
	"net/http"
)

func NewTransaction(inboxUC usecase.Inbox) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("begin new transaction")

		rsp := responder.NewNewTransaction(w)

		input, err := validation.ValidateNewTransactionRequest(r)
		if err != nil {
			slog.Info("validation error", slog.String("err", err.Error()))
			rsp.ValidationError(err)

			return
		}

		slog.Info("got valid request", slog.Any("input", input))

		transaction, err := inboxUC.MakeTransaction(
			r.Context(),
			input.UserFrom,
			input.UserTo,
			input.Amount,
			input.PaymentMethod,
			input.Currency,
			input.Description,
		)
		if err != nil {
			if errors.Is(err, usecase.ErrUserNotFound) {
				slog.Info("user not found", slog.String("err", err.Error()))

				rsp.NotFoundUser("user_from or user_to")

				return
			}

			slog.Error("error make transaction", slog.String("err", err.Error()))
			rsp.InternalError()

			return
		}

		if err = inboxUC.InspectTransaction(r.Context(), transaction); err != nil {
			slog.Info("error inspect transaction", slog.String("err", err.Error()))
			rsp.SendTransaction(transaction)

			return
		}

		if err = inboxUC.ConfirmTransaction(r.Context(), transaction); err != nil {
			slog.Error("error confirm transaction", slog.String("err", err.Error()))
			rsp.InternalError()

			return
		}

		slog.Info("transaction success confirmed")
		rsp.SendTransaction(transaction)
	}
}
