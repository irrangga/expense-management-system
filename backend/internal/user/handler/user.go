package handler

import (
	"backend/internal/user/entity"
	"backend/internal/user/handler/dto"
	"backend/internal/user/mapper"
	"backend/internal/user/usecase"
	"backend/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) *userHandler {
	return &userHandler{
		userUsecase,
	}
}

func (h *userHandler) UserLogin(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httputil.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input := entity.User{
		Email: req.Email,
	}

	user, token, err := h.userUsecase.UserLogin(ctx.Request.Context(), input)
	if err != nil {
		httputil.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToLoginResponse(token, user))
}
