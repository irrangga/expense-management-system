package usecase

import (
	"backend/internal/user/entity"
	"backend/internal/user/repo"
	"backend/pkg/token"
	"context"
)

type userUsecase struct {
	userRepo repo.UserRepo
	token    token.Token
}

func NewUserUsecase(
	userRepo repo.UserRepo,
	token token.Token,
) UserUsecase {
	return &userUsecase{
		userRepo,
		token,
	}
}

func (uc *userUsecase) UserLogin(ctx context.Context, input entity.User) (entity.User, string, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return entity.User{}, "", err
	}

	token, err := uc.token.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return entity.User{}, "", err
	}

	return user, token, nil
}
