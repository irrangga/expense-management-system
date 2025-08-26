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
		ProcessedAt: nil,
	}

	err := r.db.WithContext(ctx).Create(&expenseModel).Error
	if err != nil {
		return entity.Expense{}, err
	}

	return mapper.ToExpenseEntity(expenseModel), nil
}

func (r *expenseRepo) GetExpensesPaginated(
	ctx context.Context,
	userID int64,
	status string,
	page int,
	pageSize int,
) ([]entity.Expense, int, error) {
	var expenseModels []model.Expense
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.WithContext(ctx)

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.
		Preload("User").
		Order("submitted_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&expenseModels).Error
	if err != nil {
		return []entity.Expense{}, 0, err
	}

	query = r.db.WithContext(ctx).Model(&model.Expense{})

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err = query.Count(&total).Error
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
