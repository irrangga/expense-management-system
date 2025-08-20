package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, h *expenseHandler) {
	expense := router.Group("/expenses")

	expense.POST("", h.SubmitExpense)
}
