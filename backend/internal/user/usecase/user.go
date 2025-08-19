package usecase

import (
	"backend/internal/user/entity"
	"backend/internal/user/repo"
	"context"
)

type userUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(
	userRepo repo.UserRepo,
) UserUsecase {
	return &userUsecase{
		userRepo,
	}
}

func (uc *userUsecase) UserLogin(ctx context.Context, input entity.User) (entity.User, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
