package model

import (
	"backend/internal/user/repo/model"
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          int64
	UserID      int64
	User        model.User
	AmountIDR   int64 `gorm:"column:amount_idr"`
	Description string
	ReceiptURL  string
	Status      string
	ExternalID  *uuid.UUID
	SubmittedAt time.Time
	ProcessedAt *time.Time
}
