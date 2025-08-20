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
