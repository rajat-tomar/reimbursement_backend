package model

import (
	"time"
)

type Expense struct {
	Id          int       `json:"id,omitempty"`
	Category    string    `json:"category"`
	Amount      int       `json:"amount"`
	ExpenseDate time.Time `json:"expense_date"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
