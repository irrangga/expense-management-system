package usecase

import (
	"backend/internal/user/entity"
	"context"
)

type UserUsecase interface {
	UserLogin(ctx context.Context, input entity.User) (entity.User, string, error)
}
