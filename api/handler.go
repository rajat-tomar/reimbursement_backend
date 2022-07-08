package api

type Controllers struct {
	ExpenseController ExpenseController
	OAuthController   OAuthController
}

func NewControllers() *Controllers {
	return &Controllers{
		ExpenseController: NewExpenseController(),
		OAuthController:   NewOAuthController(),
	}
}
