package mapper

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/handler/dto"
	"backend/pkg/constant"
	"backend/pkg/httputil"
)

func ToExpenseResponse(expense entity.Expense) dto.ExpenseResponse {
	expenseResponse := dto.ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		AmountIDR:   expense.AmountIDR,
		Description: expense.Description,
		ReceiptURL:  expense.ReceiptURL,
		Status:      expense.Status,
		SubmittedAt: expense.SubmittedAt,
		ProcessedAt: expense.ProcessedAt,
	}

	if expense.AmountIDR < constant.ApprovalThreshold {
		expenseResponse.RequiresApproval = false
		expenseResponse.AutoApproved = true
	} else {
		expenseResponse.RequiresApproval = true
		expenseResponse.AutoApproved = false
	}

	return expenseResponse
}

func ToExpenseResponses(expenses []entity.Expense) []dto.ExpenseResponse {
	var expenseResponses []dto.ExpenseResponse
	for _, expense := range expenses {
		expenseResponses = append(expenseResponses, ToExpenseResponse(expense))
	}
	return expenseResponses
}

func ToExpenseResponsesPaginated(
	expenses []entity.Expense,
	page, pageSize, total int,
) httputil.PaginatedResponse[[]dto.ExpenseResponse] {
	totalPages := (total + pageSize - 1) / pageSize

	return httputil.PaginatedResponse[[]dto.ExpenseResponse]{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		Data:       ToExpenseResponses(expenses),
	}
}
