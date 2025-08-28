package usecase

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
	"backend/internal/payment"
	"backend/internal/payment/dto"
	"backend/pkg/constant"
	"context"
	"time"

	"github.com/google/uuid"
)

type expenseUsecase struct {
	expenseRepo repo.ExpenseRepo
	payment     payment.Payment
}

func NewExpenseUsecase(
	expenseRepo repo.ExpenseRepo,
	payment payment.Payment,
) ExpenseUsecase {
	return &expenseUsecase{
		expenseRepo,
		payment,
	}
}

func (uc *expenseUsecase) SubmitExpense(ctx context.Context, input entity.Expense) (entity.Expense, error) {
	if input.AmountIDR < constant.ApprovalThreshold {
		input.Status = constant.ExpenseStatusApproved

		// Process payment.
		payment, err := uc.payment.ProcessPayment(dto.PaymentRequest{
			Amount:     input.AmountIDR,
			ExternalID: uuid.NewString(),
		})
		if err != nil {
			return entity.Expense{}, err
		}

		if payment.Status == "success" {
			input.Status = constant.ExpenseStatusCompleted
		}

	} else {
		input.Status = constant.ExpenseStatusPending
	}

	input.SubmittedAt = time.Now()

	return uc.expenseRepo.SubmitExpense(ctx, input)
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

	// Process payment.
	if status == constant.ExpenseStatusApproved {
		payment, err := uc.payment.ProcessPayment(dto.PaymentRequest{
			Amount:     expense.AmountIDR,
			ExternalID: uuid.NewString(),
		})
		if err != nil {
			return entity.Expense{}, err
		}

		if payment.Status == "success" {
			expense.Status = constant.ExpenseStatusCompleted
		}
	}

	// Update expense.
	return uc.expenseRepo.UpdateExpense(ctx, expense)
}
