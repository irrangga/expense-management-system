package payment

import "backend/internal/payment/dto"

type Payment interface {
	ProcessPayment(request dto.PaymentRequest) (dto.PaymentResponse, error)
}
