package httputil

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"error": message,
	})
}

type PaginatedResponse[T any] struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       T   `json:"data"`
}
