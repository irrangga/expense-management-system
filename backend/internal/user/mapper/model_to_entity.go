package mapper

import (
	"backend/internal/user/entity"
	"backend/internal/user/repo/model"
)

func ToUserEntity(userModel model.User) entity.User {
	return entity.User{
		ID:        userModel.ID,
		Email:     userModel.Email,
		Name:      userModel.Name,
		Role:      userModel.Role,
		CreatedAt: userModel.CreatedAt,
	}
}
