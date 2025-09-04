package usecase

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
	"backend/internal/payment"
	"backend/internal/task"
	"backend/internal/task/dto"
	"backend/pkg/constant"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

type expenseUsecase struct {
	expenseRepo repo.ExpenseRepo
	payment     payment.Payment
	task        task.Task
	asynqClient *asynq.Client
}

func NewExpenseUsecase(
	expenseRepo repo.ExpenseRepo,
	payment payment.Payment,
	task task.Task,
	asynqClient *asynq.Client,
) ExpenseUsecase {
	return &expenseUsecase{
		expenseRepo,
		payment,
		task,
		asynqClient,
	}
}

func (uc *expenseUsecase) SubmitExpense(ctx context.Context, input entity.Expense) (entity.Expense, error) {
	if input.AmountIDR < constant.ApprovalThreshold {
		input.Status = constant.ExpenseStatusApproved

		externalID := uuid.New().String()
		input.ExternalID = &externalID
	} else {
		input.Status = constant.ExpenseStatusPending
	}

	input.SubmittedAt = time.Now()

	expense, err := uc.expenseRepo.SubmitExpense(ctx, input)
	if err != nil {
		return entity.Expense{}, err
	}

	// Process payment in background if approved.
	if input.Status == constant.ExpenseStatusApproved {
		err = uc.sendPaymentProcessTask(expense)
		if err != nil {
			return entity.Expense{}, err
		}
	}

	return expense, nil
}

func (uc *expenseUsecase) GetExpenses(
	ctx context.Context,
	userID int64,
	status string,
	page, pageSize int,
) ([]entity.Expense, int, error) {
	return uc.expenseRepo.GetExpensesPaginated(ctx, userID, status, page, pageSize)
}

func (uc *expenseUsecase) GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error) {
	return uc.expenseRepo.GetExpenseByID(ctx, id)
}

func (uc *expenseUsecase) ApproveExpense(ctx context.Context, id int64) (entity.Expense, error) {
	return uc.updateExpenseStatus(ctx, id, constant.ExpenseStatusApproved)
}

func (uc *expenseUsecase) RejectExpense(ctx context.Context, id int64) (entity.Expense, error) {
	return uc.updateExpenseStatus(ctx, id, constant.ExpenseStatusRejected)
}

func (uc *expenseUsecase) updateExpenseStatus(ctx context.Context, id int64, status string) (entity.Expense, error) {
	// Get expense.
	expense, err := uc.expenseRepo.GetExpenseByID(ctx, id)
	if err != nil {
		return entity.Expense{}, err
	}

	now := time.Now()
	expense.ProcessedAt = &now
	expense.Status = status

	// Process payment in background if approved.
	if status == constant.ExpenseStatusApproved {
		externalID := uuid.New().String()
		expense.ExternalID = &externalID

		err = uc.sendPaymentProcessTask(expense)
		if err != nil {
			return entity.Expense{}, err
		}
	}

	// Update expense.
	return uc.expenseRepo.UpdateExpense(ctx, expense)
}

func (uc *expenseUsecase) sendPaymentProcessTask(expense entity.Expense) error {
	if expense.ExternalID != nil {
		// Process payment in background.
		paymentTask, err := uc.task.NewPaymentProcessTask(dto.PaymentTaskPayload{
			ExpenseID:  expense.ID,
			Amount:     expense.AmountIDR,
			ExternalID: *expense.ExternalID,
		})
		if err != nil {
			return err
		}

		// Process in 5 seconds to mimic real payment processing.
		info, err := uc.asynqClient.Enqueue(paymentTask, asynq.ProcessIn(5*time.Second))
		if err != nil {
			return err
		}
		log.Printf(" [*] Successfully enqueued task: %+v", info)
	}

	return nil
}
