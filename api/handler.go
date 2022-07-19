package api

type Controllers struct {
	ExpenseController       ExpenseController
	UserController          UserController
	ReimbursementController ReimbursementController
}

func NewControllers() *Controllers {
	return &Controllers{
		ExpenseController:       NewExpenseController(),
		UserController:          NewUserController(),
		ReimbursementController: NewReimbursementController(),
	}
}
