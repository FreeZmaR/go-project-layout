package definition

import "github.com/FreeZmaR/go-service-structure/template/internal/lib/utils"

const (
	CardPaymentMethod   PaymentMethod = 1
	WalletPaymentMethod PaymentMethod = 2
	PayPalPaymentMethod PaymentMethod = 3
	BankTransferMethod  PaymentMethod = 4
)

type PaymentMethod int

func IsPaymentMethod(pm int) bool {
	return utils.OneOf(
		PaymentMethod(pm),
		CardPaymentMethod,
		WalletPaymentMethod,
		PayPalPaymentMethod,
		BankTransferMethod,
	)
}

func (pm PaymentMethod) String() string {
	switch pm {
	case CardPaymentMethod:
		return "card"
	case WalletPaymentMethod:
		return "wallet"
	case PayPalPaymentMethod:
		return "paypal"
	case BankTransferMethod:
		return "bank_transfer"
	default:
		return "unknown"
	}
}
