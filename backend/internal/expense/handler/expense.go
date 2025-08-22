package handler

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/handler/dto"
	"backend/internal/expense/mapper"
	"backend/internal/expense/usecase"
	"backend/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type expenseHandler struct {
	expenseUsecase usecase.ExpenseUsecase
}

func NewExpenseHandler(
	expenseUsecase usecase.ExpenseUsecase,
) *expenseHandler {
	return &expenseHandler{
		expenseUsecase,
	}
}

func (h *expenseHandler) SubmitExpense(ctx *gin.Context) {
	var req dto.SubmitExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input := entity.Expense{
		UserID:      ctx.GetInt64("userID"),
		AmountIDR:   req.AmountIDR,
		Description: req.Description,
		ReceiptURL:  req.ReceiptURL,
	}

	expense, err := h.expenseUsecase.SubmitExpense(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToExpenseResponse(expense))
}

func (h *expenseHandler) GetExpenseByID(ctx *gin.Context) {
	id, err := httputil.GetPathParamInt64(ctx, "id")
	if err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.expenseUsecase.GetExpenseByID(ctx.Request.Context(), id)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToExpenseResponse(expense))
}
