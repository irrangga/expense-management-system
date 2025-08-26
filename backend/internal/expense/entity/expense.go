package entity

import (
	"backend/internal/user/entity"
	"time"
)

type Expense struct {
	ID          int64
	UserID      int64
	User        entity.User
	AmountIDR   int64
	Description string
	ReceiptURL  string
	Status      string
	SubmittedAt time.Time
	ProcessedAt *time.Time
}
