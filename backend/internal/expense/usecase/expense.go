package usecase

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
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
	now := time.Now()
	input.SubmittedAt = now
	input.ProcessedAt = now

	return uc.expenseRepo.SubmitExpense(ctx, input)
}

func (uc *expenseUsecase) GetExpenses(ctx context.Context, page, pageSize int) ([]entity.Expense, int, error) {
	return uc.expenseRepo.GetExpensesPaginated(ctx, page, pageSize)
}

func (uc *expenseUsecase) GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error) {
	return uc.expenseRepo.GetExpenseByID(ctx, id)
}
