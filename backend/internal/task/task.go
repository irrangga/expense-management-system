package task

import (
	"backend/internal/expense/repo"
	"backend/internal/payment"
	paymentdto "backend/internal/payment/dto"
	"backend/internal/task/dto"
	"backend/pkg/constant"
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

type task struct {
	expenseRepo repo.ExpenseRepo
	payment     payment.Payment
}

func NewTask(
	expenseRepo repo.ExpenseRepo,
	payment payment.Payment,
) Task {
	return &task{
		expenseRepo,
		payment,
	}
}

func (t *task) NewPaymentProcessTask(payment dto.PaymentTaskPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(payment)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(constant.PaymentType, payload), nil
}

func (t *task) HandlePaymentProcessTask(ctx context.Context, a *asynq.Task) error {
	var p dto.PaymentTaskPayload
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return err
	}

	// Get expense.
	expense, err := t.expenseRepo.GetExpenseByID(ctx, p.ExpenseID)
	if err != nil {
		return err
	}

	// Process payment.
	payment, err := t.payment.ProcessPayment(paymentdto.PaymentRequest{
		Amount:     p.Amount,
		ExternalID: p.ExternalID,
	})
	if err != nil {
		return err
	}

	if payment.Status == "success" {
		expense.Status = constant.ExpenseStatusCompleted
	}

	// Update expense.
	_, err = t.expenseRepo.UpdateExpense(ctx, expense)
	if err != nil {
		return err
	}

	log.Printf(" [*] Process Payment for Expense %d", p.ExpenseID)
	return nil
}
