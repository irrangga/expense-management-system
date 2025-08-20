package model

import "time"

type Expense struct {
	ID          int64
	UserID      int64
	AmountIDR   int64 `gorm:"column:amount_idr"`
	Description string
	ReceiptURL  string
	Status      string
	SubmittedAt time.Time
	ProcessedAt time.Time
}
