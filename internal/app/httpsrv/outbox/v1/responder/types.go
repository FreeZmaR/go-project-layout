package responder

type TransactionUser struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	AvailablePaymentMethods []string `json:"available_payment_methods"`
}

type Transaction struct {
	ID            string  `json:"id"`
	Amount        int     `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	ErrorCode     *string `json:"error_code,omitempty"`
	ErrorDetails  *string `json:"error_details,omitempty"`
}

type TransactionHistory struct {
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}
