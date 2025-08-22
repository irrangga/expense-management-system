package usecase

import (
	"backend/generated/mockgen/expense"
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_expenseUsecase_SubmitExpense(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := expense.NewMockExpenseRepo(ctrl)

	expense := entity.Expense{
		AmountIDR:   10000,
		Description: "Expense 1",
		ReceiptURL:  "https://example.com/receipt.jpg",
	}

	type fields struct {
		expenseRepo repo.ExpenseRepo
	}
	type args struct {
		ctx   context.Context
		input entity.Expense
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Expense
		wantErr error
	}{
		{
			name: "SubmitExpense returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: expense,
			},
			mock: func() {
				expenseRepoMock.EXPECT().SubmitExpense(gomock.Any(), gomock.Any()).Return(expense, nil)
			},
			want:    expense,
			wantErr: nil,
		},
		{
			name: "SubmitExpense returns error",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: expense,
			},
			mock: func() {
				expenseRepoMock.EXPECT().SubmitExpense(gomock.Any(), gomock.Any()).Return(entity.Expense{}, assert.AnError)
			},
			want:    entity.Expense{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &expenseUsecase{
				expenseRepo: tt.fields.expenseRepo,
			}
			tt.mock()

			got, err := uc.SubmitExpense(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_expenseUsecase_GetExpenseByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := expense.NewMockExpenseRepo(ctrl)

	expense := entity.Expense{
		ID:          1,
		AmountIDR:   10000,
		Description: "Expense 1",
		ReceiptURL:  "https://example.com/receipt.jpg",
	}

	type fields struct {
		expenseRepo repo.ExpenseRepo
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Expense
		wantErr error
	}{
		{
			name: "GetExpenseByID returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpenseByID(gomock.Any(), int64(1)).Return(expense, nil)
			},
			want:    expense,
			wantErr: nil,
		},
		{
			name: "GetExpenseByID returns error",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  int64(2),
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpenseByID(gomock.Any(), int64(2)).Return(entity.Expense{}, assert.AnError)
			},
			want:    entity.Expense{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &expenseUsecase{
				expenseRepo: tt.fields.expenseRepo,
			}
			tt.mock()

			got, err := uc.GetExpenseByID(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
