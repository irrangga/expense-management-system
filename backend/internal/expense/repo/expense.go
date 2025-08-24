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

func (r *expenseRepo) GetExpensesPaginated(
	ctx context.Context,
	page int,
	pageSize int,
) ([]entity.Expense, int, error) {
	var expenseModels []model.Expense
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).
		Limit(pageSize).
		Offset(offset).
		Find(&expenseModels).Error
	if err != nil {
		return []entity.Expense{}, 0, err
	}

	err = r.db.WithContext(ctx).Model(&model.Expense{}).Count(&total).Error
	if err != nil {
		return []entity.Expense{}, 0, err
	}

	return mapper.ToExpenseEntities(expenseModels), int(total), err
}

func (r *expenseRepo) GetExpenseByID(ctx context.Context, id int64) (entity.Expense, error) {
	var expenseModel model.Expense

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&expenseModel).Error
	if err != nil {
		return entity.Expense{}, err
	}

	return mapper.ToExpenseEntity(expenseModel), nil
}
