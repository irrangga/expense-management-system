package usecase

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
	"backend/pkg/constant"
	"context"
	"time"
)

type expenseUsecase struct {
	expenseRepo repo.ExpenseRepo
}

func NewExpenseUsecase(
	expenseRepo repo.ExpenseRepo,
) ExpenseUsecase {
	return &expenseUsecase{
		expenseRepo,
	}
}

func (uc *expenseUsecase) SubmitExpense(ctx context.Context, input entity.Expense) (entity.Expense, error) {
	if input.AmountIDR < constant.ApprovalThreshold {
		input.Status = constant.ExpenseStatusApproved
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
