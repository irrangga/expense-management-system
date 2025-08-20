package mapper

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/handler/dto"
)

func ToExpenseResponse(expense entity.Expense) dto.ExpenseResponse {
	return dto.ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		AmountIDR:   expense.AmountIDR,
		Description: expense.Description,
		ReceiptURL:  expense.ReceiptURL,
		Status:      expense.Status,
		SubmittedAt: expense.SubmittedAt,
		ProcessedAt: expense.ProcessedAt,
	}
}
