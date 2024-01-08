package responder

import (
	"github.com/FreeZmaR/go-project-layout/internal/domain/aggregate"
	"github.com/FreeZmaR/go-project-layout/internal/domain/definition"
	"github.com/FreeZmaR/go-project-layout/internal/lib/responder"
	"github.com/FreeZmaR/go-project-layout/internal/lib/utils"
	"log/slog"
	"net/http"
)

type NewTransaction struct {
	base *responder.HTTP
}

func NewNewTransaction(w http.ResponseWriter) *NewTransaction {
	return &NewTransaction{base: responder.NewHTTP(w)}
}

func (r *NewTransaction) ValidationError(err error) {
	r.sendError(
		http.StatusBadRequest,
		definition.ValidationFailedErrCode.String(),
		err.Error(),
	)
}

func (r *NewTransaction) InternalError() {
	r.sendError(
		http.StatusInternalServerError,
		definition.InternalErrCode.String(),
		definition.InternalErrCode.Description(),
	)
}

func (r *NewTransaction) SendTransaction(transaction *aggregate.Transaction) {
	type response struct {
		ID            string  `json:"id"`
		Amount        int     `json:"amount"`
		Currency      string  `json:"currency"`
		PaymentMethod string  `json:"payment_method"`
		Status        string  `json:"status"`
		ErrorCode     *string `json:"error_code,omitempty"`
		ErrorDetails  *string `json:"error_details,omitempty"`
	}

	var errCode *string
	if transaction.ErrorCode() != nil {
		errCode = utils.WithPtr(transaction.ErrorCode().String())
	}

	r.send(
		http.StatusOK,
		response{
			ID:            transaction.ID().String(),
			Amount:        transaction.Amount(),
			Currency:      transaction.Currency().String(),
			PaymentMethod: transaction.PaymentMethod().String(),
			Status:        transaction.Status().String(),
			ErrorCode:     errCode,
			ErrorDetails:  transaction.ErrorDescription(),
		},
	)
}

func (r *NewTransaction) NotFoundUser(details string) {
	r.sendError(
		http.StatusBadRequest,
		definition.ValidationFailedErrCode.String(),
		details,
	)
}

func (r *NewTransaction) sendError(httpCode int, code, details string) {
	type errorResponse struct {
		Code    string `json:"code"`
		Details string `json:"details"`
	}

	r.send(httpCode, errorResponse{Code: code, Details: details})
}

func (r *NewTransaction) send(code int, body any) {
	if err := r.base.SendJSON(code, body); err != nil {
		slog.Error("error send response", slog.String("err", err.Error()))

		return
	}

	slog.Info("response sent", slog.Int("code", code), slog.Any("body", body))
}
