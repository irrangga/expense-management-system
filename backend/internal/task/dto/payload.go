package dto

type PaymentTaskPayload struct {
	ExpenseID  int64
	Amount     int64
	ExternalID string
}
