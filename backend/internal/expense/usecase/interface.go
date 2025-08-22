package usecase

import (
	"backend/internal/expense/entity"
	"context"
)

type ExpenseUsecase interface {
	SubmitExpense(ctx context.Context, input entity.Expense) (entity.Expense, error)
	GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error)
}
