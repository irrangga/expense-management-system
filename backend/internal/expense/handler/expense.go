package handler

import (
	"backend/internal/expense/entity"
	"backend/internal/expense/handler/dto"
	"backend/internal/expense/mapper"
	"backend/internal/expense/usecase"
	"backend/pkg/constant"
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

	if req.AmountIDR < constant.MinExpenseAmount {
		httputil.Error(ctx, http.StatusBadRequest, "Amount must be greater than Rp 10.000")
		return
	}
	if req.AmountIDR > constant.MaxExpenseAmount {
		httputil.Error(ctx, http.StatusBadRequest, "Amount must be less than Rp 50.000.000")
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

func (h *expenseHandler) GetExpenses(ctx *gin.Context) {
	page, _ := httputil.GetQueryParamInt(ctx, "page")
	pageSize, _ := httputil.GetQueryParamInt(ctx, "page_size")
	status := ctx.Query("status")

	var userID int64 = 0
	if ctx.GetString("role") == constant.UserRoleEmployee {
		userID = ctx.GetInt64("userID")
	}

	// Default values.
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	expenses, total, err := h.expenseUsecase.GetExpenses(ctx.Request.Context(), userID, status, page, pageSize)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToExpenseResponsesPaginated(expenses, page, pageSize, total))
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
