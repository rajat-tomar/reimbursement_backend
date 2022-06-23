package api

type Controllers struct {
	ExpenseController ExpenseController
}

func NewControllers() *Controllers {
	return &Controllers{
		ExpenseController: NewExpenseController(),
	}
}
