package repo

import (
	"backend/internal/expense/entity"
	"context"
)

type ExpenseRepo interface {
	SubmitExpense(ctx context.Context, expense entity.Expense) (entity.Expense, error)
	GetExpensesPaginated(ctx context.Context, userID int64, status string, page int, pageSize int) ([]entity.Expense, int, error)
	GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error)
}
