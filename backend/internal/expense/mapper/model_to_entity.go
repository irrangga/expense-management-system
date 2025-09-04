package mapper

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/repo/model"
	userentity "backend/internal/user/entity"
)

func ToExpenseEntity(expenseModel model.Expense) entity.Expense {
	var externalIDString *string
	if expenseModel.ExternalID != nil {
		s := expenseModel.ExternalID.String()
		externalIDString = &s
	}

	return entity.Expense{
		ID:     expenseModel.ID,
		UserID: expenseModel.UserID,
		User: userentity.User{
			ID:    expenseModel.User.ID,
			Email: expenseModel.User.Email,
			Name:  expenseModel.User.Name,
			Role:  expenseModel.User.Role,
		},
		AmountIDR:   expenseModel.AmountIDR,
		Description: expenseModel.Description,
		ReceiptURL:  expenseModel.ReceiptURL,
		Status:      expenseModel.Status,
		ExternalID:  externalIDString,
		SubmittedAt: expenseModel.SubmittedAt,
		ProcessedAt: expenseModel.ProcessedAt,
	}
}

func ToExpenseEntities(expenseModels []model.Expense) []entity.Expense {
	var expenses []entity.Expense
	for _, expenseModel := range expenseModels {
		expenses = append(expenses, ToExpenseEntity(expenseModel))
	}
	return expenses
}
