package responder

import (
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/aggregate"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/definition"
	"github.com/FreeZmaR/go-service-structure/template/internal/domain/objvalue"
	"github.com/FreeZmaR/go-service-structure/template/internal/lib/responder"
	"log/slog"
	"net/http"
)

type GetTransaction struct {
	base *responder.HTTP
}

func NewShowTransaction(w http.ResponseWriter) *GetTransaction {
	return &GetTransaction{base: responder.NewHTTP(w)}
}

func (r *GetTransaction) ValidationError(err error) {
	r.sendError(
		http.StatusBadRequest,
		definition.ValidationFailedErrCode.String(),
		err.Error(),
	)
}

func (r *GetTransaction) InternalError() {
	r.sendError(
		http.StatusInternalServerError,
		definition.InternalErrCode.String(),
		definition.InternalErrCode.Description(),
	)
}

func (r *GetTransaction) SendTransaction(transaction *aggregate.Transaction) {
	type response struct {
		UserFrom    TransactionUser     `json:"user_from"`
		UserTo      TransactionUser     `json:"user_to"`
		Transaction Transaction         `json:"transaction"`
		History     *TransactionHistory `json:"history,omitempty"`
	}

	var history *TransactionHistory

	if transaction.History().OldStatus != objvalue.TransactionStatusUnknown {
		history = &TransactionHistory{
			OldStatus: transaction.History().OldStatus.String(),
			NewStatus: transaction.History().NewStatus.String(),
		}
	}

	availablePMFrom := make([]string, 0, len(transaction.UserFrom().AvailablePaymentMethods))
	for _, pm := range transaction.UserFrom().AvailablePaymentMethods {
		availablePMFrom = append(availablePMFrom, pm.String())
	}

	availablePMTo := make([]string, 0, len(transaction.UserTo().AvailablePaymentMethods))
	for _, pm := range transaction.UserTo().AvailablePaymentMethods {
		availablePMTo = append(availablePMTo, pm.String())
	}

	var errCodee *string

	r.send(
		http.StatusOK,
		response{
			UserFrom: TransactionUser{
				ID:                      transaction.UserFrom().ID.String(),
				Name:                    transaction.UserFrom().Name,
				AvailablePaymentMethods: availablePMFrom,
			},
			UserTo: TransactionUser{
				ID:                      transaction.UserTo().ID.String(),
				Name:                    transaction.UserTo().Name,
				AvailablePaymentMethods: availablePMTo,
			},
			Transaction: Transaction{
				ID:            transaction.ID().String(),
				Amount:        transaction.Amount(),
				Currency:      transaction.Currency().String(),
				PaymentMethod: transaction.PaymentMethod().String(),
				Status:        transaction.Status().String(),
				ErrorCode:     errCodee,
				ErrorDetails:  transaction.ErrorDescription(),
			},
			History: history,
		},
	)
}

func (r *GetTransaction) NotFound() {
	r.sendError(
		http.StatusNotFound,
		definition.NotFoundErrCode.String(),
		definition.NotFoundErrCode.Description(),
	)
}

func (r *GetTransaction) sendError(httpCode int, code, details string) {
	type errorResponse struct {
		Code    string `json:"code"`
		Details string `json:"details"`
	}

	r.send(httpCode, errorResponse{Code: code, Details: details})
}

func (r *GetTransaction) send(code int, body any) {
	if err := r.base.SendJSON(code, body); err != nil {
		slog.Error("error send response", slog.String("err", err.Error()))

		return
	}

	slog.Info("response sent", slog.Int("code", code), slog.Any("body", body))
}
