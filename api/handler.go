package api

type Controllers struct {
	ExpenseController ExpenseController
	UserController    UserController
}

func NewControllers() *Controllers {
	return &Controllers{
		ExpenseController: NewExpenseController(),
		UserController:    NewUserController(),
	}
}
