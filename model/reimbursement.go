package model

type Reimbursement struct {
	Id          int    `json:"id"`
	Amount      int    `json:"amount"`
	UserId      int    `json:"user_id"`
	ExpenseId   int    `json:"expense_id"`
	Status      string `json:"status"`
	ProcessedOn string `json:"processed_on"`
}
