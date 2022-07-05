package model

import (
	"time"
)

type Expense struct {
	Id          int       `json:"id,omitempty"`
	Amount      int       `json:"amount"`
	ExpenseDate time.Time `json:"expense_date"`
	CreatedAt   time.Time `json:"created_at"`
}
