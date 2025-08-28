package constant

const (
	UserRoleEmployee = "employee"
	UserRoleManager  = "manager"
)

const (
	ExpenseStatusPending   = "pending"
	ExpenseStatusApproved  = "approved"
	ExpenseStatusRejected  = "rejected"
	ExpenseStatusCompleted = "completed"
)

const (
	MinExpenseAmount  = 10_000
	MaxExpenseAmount  = 50_000_000
	ApprovalThreshold = 1_000_000
)
