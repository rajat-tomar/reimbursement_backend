package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
	"testing"
)

func TestRepositoryExpense_GetById(t *testing.T) {
	sqlStatement := `INSERT INTO expenses(Id, Amount) VALUES($1, $2)`
	_ = testDb.QueryRow(sqlStatement, 9, 100)

	row1 := testDb.QueryRow(`select id, amount from expenses where id = 9`)
	var expense model.Expense
	err := row1.Scan(
		&expense.Id,
		&expense.Amount,
	)
	fmt.Println(expense.Id, expense.Amount)
	if err != nil {
		config.Logger.Panicw("cannot connect to the database", "error", err)
		panic(err)
	}

	expenseRepository := expenseRepository{db: testDb}
	Expense, err := expenseRepository.GetById(9)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, Expense)
	assert.Equal(t, Expense.Id, int64(9))
	assert.Equal(t, Expense.Amount, int64(100))
}
