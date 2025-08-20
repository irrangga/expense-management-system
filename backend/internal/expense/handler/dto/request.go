package dto

type SubmitExpenseRequest struct {
	UserID      int64  `json:"user_id" binding:"required"`
	AmountIDR   int64  `json:"amount_idr" binding:"required"`
	Description string `json:"description" binding:"required"`
	ReceiptURL  string `json:"receipt_url"`
}
