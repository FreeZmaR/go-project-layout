package definition

const (
	ValidationFailedErrCode  ErrorCode = "A101"
	InsufficientFundsErrCode ErrorCode = "A102"
	InternalErrCode          ErrorCode = "A103"
	NotFoundErrCode          ErrorCode = "A104"
)

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) Description() string {
	switch e {
	case ValidationFailedErrCode:
		return "Validation failed"
	case InsufficientFundsErrCode:
		return "Insufficient funds"
	case InternalErrCode:
		return "Internal error"
	case NotFoundErrCode:
		return "Not found"
	default:
		return "Unknown error"
	}
}
