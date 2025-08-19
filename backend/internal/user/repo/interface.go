package repo

import (
	"backend/internal/user/entity"
	"context"
)

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
