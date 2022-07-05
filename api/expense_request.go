package api

type ExpenseRequest struct {
	Id          int    `json:"id,omitempty"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	ExpenseDate string `json:"expense_date"`
}
