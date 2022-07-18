package request_model

type ExpenseRequest struct {
	Id          string `json:"id"`
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	ExpenseDate string `json:"expense_date"`
	Status      string `json:"status"`
	UserId      int    `json:"user_id"`
}
