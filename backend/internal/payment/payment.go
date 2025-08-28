package payment

import (
	"backend/config"
	"backend/internal/payment/dto"
	"backend/pkg/httputil"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type payment struct {
	httpClient httputil.Client
}

func NewPayment(
	httpClient httputil.Client,
) Payment {
	return &payment{
		httpClient,
	}
}

func (p *payment) ProcessPayment(request dto.PaymentRequest) (dto.PaymentResponse, error) {
	respBody, respStatusCode, err := p.httpClient.SendRequest(
		"POST",
		fmt.Sprintf("%+v/v1/payments", config.Cfg.Payment.PaymentProcessorURL),
		map[string]string{},
		request,
	)
	if err != nil || respStatusCode != http.StatusOK {
		return dto.PaymentResponse{}, err
	}

	var result httputil.Response[dto.PaymentResponse]
	if err = json.Unmarshal(respBody, &result); err != nil {
		return dto.PaymentResponse{}, err
	}

	if result.Message != "" {
		return dto.PaymentResponse{}, errors.New(result.Message)
	}

	return result.Data, nil
}
