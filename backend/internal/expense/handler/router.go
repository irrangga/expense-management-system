package handler

import (
	"backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h *expenseHandler) {
	expense := router.Group("/expenses")

	expense.POST("", middleware.AuthMiddleware(), h.SubmitExpense)
	expense.GET("", middleware.AuthMiddleware(), h.GetExpenses)
	expense.GET("/:id", middleware.AuthMiddleware(), h.GetExpenseByID)
}
