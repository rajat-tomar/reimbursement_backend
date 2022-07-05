package api

type ExpenseRequest struct {
	Id          int    `json:"id,omitempty"`
	Amount      int    `json:"amount"`
	ExpenseDate string `json:"expense_date"`
}
