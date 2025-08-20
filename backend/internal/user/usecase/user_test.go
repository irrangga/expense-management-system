package usecase

import (
	"backend/generated/mockgen/pkg"
	"backend/generated/mockgen/user"
	"backend/internal/user/entity"
	"backend/internal/user/repo"
	"backend/pkg/token"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_userUsecase_UserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	userRepoMock := user.NewMockUserRepo(ctrl)
	tokenMock := pkg.NewMockToken(ctrl)

	user := entity.User{
		ID:    1,
		Email: "email",
		Role:  "employee",
	}

	type fields struct {
		userRepo repo.UserRepo
		token    token.Token
	}
	type args struct {
		ctx   context.Context
		input entity.User
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mock      func()
		wantUser  entity.User
		wantToken string
		wantErr   error
	}{
		{
			name: "UserLogin returns user successfully",
			fields: fields{
				userRepo: userRepoMock,
				token:    tokenMock,
			},
			args: args{
				ctx: context.Background(),
				input: entity.User{
					Email: "email",
				},
			},
			mock: func() {
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), "email").Return(user, nil)
				tokenMock.EXPECT().GenerateJWT(user.ID, user.Email, user.Role).Return("token", nil)
			},
			wantUser:  user,
			wantToken: "token",
			wantErr:   nil,
		},
		{
			name: "UserLogin returns GetUserByEmail error",
			fields: fields{
				userRepo: userRepoMock,
				token:    tokenMock,
			},
			args: args{
				ctx: context.Background(),
				input: entity.User{
					Email: "email",
				},
			},
			mock: func() {
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), "email").Return(entity.User{}, assert.AnError)
			},
			wantUser:  entity.User{},
			wantToken: "",
			wantErr:   assert.AnError,
		},
		{
			name: "UserLogin returns GenerateJWT error",
			fields: fields{
				userRepo: userRepoMock,
				token:    tokenMock,
			},
			args: args{
				ctx: context.Background(),
				input: entity.User{
					Email: "email",
				},
			},
			mock: func() {
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), "email").Return(user, nil)
				tokenMock.EXPECT().GenerateJWT(user.ID, user.Email, user.Role).Return("", assert.AnError)
			},
			wantUser:  entity.User{},
			wantToken: "",
			wantErr:   assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userUsecase{
				userRepo: tt.fields.userRepo,
				token:    tt.fields.token,
			}
			tt.mock()

			gotUser, gotToken, err := uc.UserLogin(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.wantUser, gotUser)
			assert.Equal(t, tt.wantToken, gotToken)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
