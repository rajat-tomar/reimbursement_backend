package api

type Controllers struct {
	ExpenseController ExpenseController
	OAuthController   OAuthController
	UserController    UserController
}

func NewControllers() *Controllers {
	return &Controllers{
		ExpenseController: NewExpenseController(),
		OAuthController:   NewOAuthController(),
		UserController:    NewUserController(),
	}
}
