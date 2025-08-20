package dto

type SubmitExpenseRequest struct {
	AmountIDR   int64  `json:"amount_idr" binding:"required"`
	Description string `json:"description" binding:"required"`
	ReceiptURL  string `json:"receipt_url"`
}
