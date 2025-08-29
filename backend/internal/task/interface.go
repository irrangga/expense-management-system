package task

import (
	"backend/internal/task/dto"
	"context"

	"github.com/hibiken/asynq"
)

type Task interface {
	NewPaymentProcessTask(payment dto.PaymentTaskPayload) (*asynq.Task, error)
	HandlePaymentProcessTask(ctx context.Context, t *asynq.Task) error
}
