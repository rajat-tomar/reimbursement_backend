package request_model

type ExpenseRequest struct {
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	ExpenseDate string `json:"expense_date"`
}
