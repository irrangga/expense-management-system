package repo

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/mapper"
	"backend/internal/expense/repo/model"
	"context"

	"gorm.io/gorm"
)

type expenseRepo struct {
	db *gorm.DB
}

func NewExpenseRepo(
	db *gorm.DB,
) ExpenseRepo {
	return &expenseRepo{
		db,
	}
}

func (r *expenseRepo) SubmitExpense(ctx context.Context, expense entity.Expense) (entity.Expense, error) {
	expenseModel := model.Expense{
		UserID:      expense.UserID,
		AmountIDR:   expense.AmountIDR,
		Description: expense.Description,
		ReceiptURL:  expense.ReceiptURL,
		Status:      expense.Status,
		SubmittedAt: expense.SubmittedAt,
		ProcessedAt: expense.ProcessedAt,
	}

	err := r.db.WithContext(ctx).Create(&expenseModel).Error
	if err != nil {
		return entity.Expense{}, err
	}

	return mapper.ToExpenseEntity(expenseModel), nil
}

func (r *expenseRepo) GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error) {
	var expenseModel model.Expense

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&expenseModel).Error
	if err != nil {
		return entity.Expense{}, err
	}

	return mapper.ToExpenseEntity(expenseModel), nil
}
