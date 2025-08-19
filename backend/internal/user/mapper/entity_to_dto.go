package mapper

import (
	"backend/internal/user/entity"
	"backend/internal/user/handler/dto"
)

func ToLoginResponse(token string, user entity.User) dto.LoginResponse {
	return dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Role:  user.Role,
		},
	}
}
