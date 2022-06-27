package model

import (
	"time"
)

type Expense struct {
	Id        int       `json:"id,omitempty"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
