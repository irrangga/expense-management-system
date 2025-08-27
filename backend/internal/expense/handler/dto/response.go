package dto

import "time"

type ExpenseResponse struct {
	ID               int64      `json:"id"`
	UserName         string     `json:"user_name,omitempty"`
	AmountIDR        int64      `json:"amount_idr"`
	Description      string     `json:"description"`
	ReceiptURL       string     `json:"receipt_url"`
	Status           string     `json:"status"`
	RequiresApproval bool       `json:"requires_approval"`
	AutoApproved     bool       `json:"auto_approved"`
	SubmittedAt      time.Time  `json:"submitted_at"`
	ProcessedAt      *time.Time `json:"processed_at,omitempty"`
}
