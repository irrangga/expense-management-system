package mapper

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo/model"
)

func ToExpenseEntity(expenseModel model.Expense) entity.Expense {
	return entity.Expense{
		ID:          expenseModel.ID,
		UserID:      expenseModel.UserID,
		AmountIDR:   expenseModel.AmountIDR,
		Description: expenseModel.Description,
		ReceiptURL:  expenseModel.ReceiptURL,
		Status:      expenseModel.Status,
		SubmittedAt: expenseModel.SubmittedAt,
		ProcessedAt: expenseModel.ProcessedAt,
	}
}
