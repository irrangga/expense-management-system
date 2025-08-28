package dto

type PaymentRequest struct {
	Amount     int64  `json:"amount" binding:"required"`
	ExternalID string `json:"external_id" binding:"required"`
}
