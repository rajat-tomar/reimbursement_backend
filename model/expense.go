package model

import (
	"fmt"
	"time"
)

type Expense struct {
	ID        int
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Expense) String() string {
	return fmt.Sprintf("Id %s \n Amount %f", e.ID, e.Amount)
}
