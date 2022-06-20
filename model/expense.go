package model

import (
	"time"
)

type Expense struct {
	Id        int64     `json:"id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
