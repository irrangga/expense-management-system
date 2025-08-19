package repo

import (
	"backend/internal/user/entity"
	"backend/internal/user/mapper"
	"backend/internal/user/repo/model"
	"context"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(
	db *gorm.DB,
) UserRepo {
	return &userRepo{
		db,
	}
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var userModel model.User

	err := r.db.WithContext(ctx).First(&userModel, "email = ?", email).Error
	if err != nil {
		return entity.User{}, err
	}

	return mapper.ToUserEntity(userModel), nil
}
