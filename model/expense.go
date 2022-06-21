package model

import (
	"time"
)

type Expense struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
