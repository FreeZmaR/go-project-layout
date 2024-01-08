package objvalue

import "github.com/FreeZmaR/go-service-structure/template/internal/lib/utils"

const (
	TransactionStatusUnknown TransactionStatus = 0
	TransactionStatusNew     TransactionStatus = 1
	TransactionStatusSuccess TransactionStatus = 2
	TransactionStatusFail    TransactionStatus = 3
	TransactionStatusPending TransactionStatus = 4
)

type TransactionStatus int

func IsTransactionStatus(ts int) bool {
	return utils.OneOf(
		TransactionStatus(ts),
		TransactionStatusNew,
		TransactionStatusSuccess,
		TransactionStatusFail,
		TransactionStatusPending,
	)
}

func (ts TransactionStatus) String() string {
	switch ts {
	case TransactionStatusNew:
		return "new"
	case TransactionStatusSuccess:
		return "success"
	case TransactionStatusFail:
		return "fail"
	case TransactionStatusPending:
		return "pending"
	default:
		return "unknown"
	}
}
