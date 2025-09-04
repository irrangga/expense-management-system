package usecase

import (
	mock_asynq "backend/generated/mockgen/asynq"
	mock_expense "backend/generated/mockgen/expense"
	mock_task "backend/generated/mockgen/task"
	"backend/internal/expense/entity"
	"backend/internal/expense/repo"
	"backend/internal/task"
	"backend/internal/task/dto"
	"backend/pkg/asynq/client"
	"backend/pkg/constant"
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_expenseUsecase_SubmitExpense(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := mock_expense.NewMockExpenseRepo(ctrl)

	expenseBelowThreshold := entity.Expense{
		AmountIDR:   10000,
		Description: "Expense 1",
		ReceiptURL:  "https://example.com/receipt.jpg",
	}

	expenseAboveThreshold := entity.Expense{
		AmountIDR:   50000000,
		Description: "Expense 2",
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
			name: "SubmitExpense below approval threshold returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: expenseBelowThreshold,
			},
			mock: func() {
				expenseRepoMock.EXPECT().SubmitExpense(gomock.Any(), gomock.Any()).Return(expenseBelowThreshold, nil)
			},
			want:    expenseBelowThreshold,
			wantErr: nil,
		},
		{
			name: "SubmitExpense above approval threshold returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: expenseAboveThreshold,
			},
			mock: func() {
				expenseRepoMock.EXPECT().SubmitExpense(gomock.Any(), gomock.Any()).Return(expenseAboveThreshold, nil)
			},
			want:    expenseAboveThreshold,
			wantErr: nil,
		},
		{
			name: "SubmitExpense returns error",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:   context.Background(),
				input: expenseBelowThreshold,
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

	expenseRepoMock := mock_expense.NewMockExpenseRepo(ctrl)

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

func Test_expenseUsecase_GetExpenses(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := mock_expense.NewMockExpenseRepo(ctrl)

	expenses := []entity.Expense{
		{
			ID:          1,
			AmountIDR:   10000,
			Description: "Expense 1",
			ReceiptURL:  "https://example.com/receipt.jpg",
		},
		{
			ID:          2,
			AmountIDR:   20000,
			Description: "Expense 2",
			ReceiptURL:  "https://example.com/receipt.jpg",
		},
	}

	type fields struct {
		expenseRepo repo.ExpenseRepo
	}
	type args struct {
		ctx      context.Context
		userID   int64
		status   string
		page     int
		pageSize int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mock         func()
		wantExpenses []entity.Expense
		wantTotal    int
		wantErr      error
	}{
		{
			name: "GetExpenses returns expenses successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:      context.Background(),
				userID:   1,
				status:   constant.ExpenseStatusPending,
				page:     1,
				pageSize: 10,
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpensesPaginated(gomock.Any(), int64(1), constant.ExpenseStatusPending, 1, 10).Return(expenses, 2, nil)
			},
			wantExpenses: expenses,
			wantTotal:    2,
			wantErr:      nil,
		},
		{
			name: "GetExpenses returns error",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx:      context.Background(),
				userID:   1,
				status:   constant.ExpenseStatusPending,
				page:     1,
				pageSize: 10,
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpensesPaginated(gomock.Any(), int64(1), constant.ExpenseStatusPending, 1, 10).
					Return([]entity.Expense{}, 0, assert.AnError)
			},
			wantExpenses: []entity.Expense{},
			wantTotal:    0,
			wantErr:      assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &expenseUsecase{
				expenseRepo: tt.fields.expenseRepo,
			}
			tt.mock()

			gotExpenses, gotTotal, err := uc.GetExpenses(tt.args.ctx, tt.args.userID, tt.args.status, tt.args.page, tt.args.pageSize)
			assert.Equal(t, tt.wantExpenses, gotExpenses)
			assert.Equal(t, tt.wantTotal, gotTotal)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_expenseUsecase_ApproveExpense(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := mock_expense.NewMockExpenseRepo(ctrl)
	taskMock := mock_task.NewMockTask(ctrl)
	asynqClientMock := mock_asynq.NewMockAsynqClient(ctrl)

	payload, _ := json.Marshal(dto.PaymentTaskPayload{
		ExpenseID:  1,
		Amount:     150000,
		ExternalID: uuid.New().String(),
	})

	taskInfo := asynq.TaskInfo{}

	type fields struct {
		expenseRepo repo.ExpenseRepo
		task        task.Task
		asynqClient client.AsynqClient
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
			name: "ApproveExpense returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
				task:        taskMock,
				asynqClient: asynqClientMock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpenseByID(gomock.Any(), int64(1)).Return(entity.Expense{}, nil)
				taskMock.EXPECT().NewPaymentProcessTask(gomock.Any()).Return(asynq.NewTask(constant.PaymentType, payload), nil)
				asynqClientMock.EXPECT().Enqueue(gomock.Any(), gomock.Any()).Return(&taskInfo, nil)
				expenseRepoMock.EXPECT().UpdateExpense(gomock.Any(), gomock.Any()).Return(entity.Expense{}, nil)
			},
			want:    entity.Expense{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &expenseUsecase{
				expenseRepo: tt.fields.expenseRepo,
				task:        tt.fields.task,
				asynqClient: tt.fields.asynqClient,
			}
			tt.mock()

			got, err := uc.ApproveExpense(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func Test_expenseUsecase_RejectExpense(t *testing.T) {
	ctrl := gomock.NewController(t)

	expenseRepoMock := mock_expense.NewMockExpenseRepo(ctrl)

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
			name: "RejectExpense returns expense successfully",
			fields: fields{
				expenseRepo: expenseRepoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func() {
				expenseRepoMock.EXPECT().GetExpenseByID(gomock.Any(), int64(1)).Return(entity.Expense{}, nil)
				expenseRepoMock.EXPECT().UpdateExpense(gomock.Any(), gomock.Any()).Return(entity.Expense{}, nil)
			},
			want:    entity.Expense{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &expenseUsecase{
				expenseRepo: tt.fields.expenseRepo,
			}
			tt.mock()

			got, err := uc.RejectExpense(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
