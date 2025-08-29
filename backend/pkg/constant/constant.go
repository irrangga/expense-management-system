package constant

// User.
const (
	UserRoleEmployee = "employee"
	UserRoleManager  = "manager"
)

// Expense.
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

// Payment.
const (
	PaymentType          = "payment"
	PaymentStatusSuccess = "success"
)
