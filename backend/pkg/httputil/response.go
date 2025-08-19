package httputil

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"error": message,
	})
}
