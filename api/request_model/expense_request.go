package request_model

type ExpenseRequest struct {
	Id          string `json:"id,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Category    string `json:"category,omitempty"`
	ExpenseDate string `json:"expense_date,omitempty"`
	UserId      int    `json:"user_id,omitempty"`
	Status      string `json:"status,omitempty"`
}
