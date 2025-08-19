package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, h *userHandler) {
	router.POST("/auth/login", h.UserLogin)
}
