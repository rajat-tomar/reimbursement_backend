package repository

import (
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/model"
	"testing"
)

func TestGetExpenseById(t *testing.T) {
	sqlStatement := `INSERT INTO expenses(Id, Amount) VALUES($1, $2)`
	_ = testDb.QueryRow(sqlStatement, 9, 100)

	var expense model.Expense
	err := testDb.QueryRow(`select id, amount from expenses where id = 9`).Scan(&expense.Id, &expense.Amount)
	expenseRepository := expenseRepository{db: testDb}
	expenseGot, err := expenseRepository.GetExpenseById(9)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, expenseGot)
	assert.Equal(t, expenseGot.Id, 9)
	assert.Equal(t, expenseGot.Amount, 100)
}

func TestCreateExpense(t *testing.T) {
	expenseExpected := model.Expense{
		Amount: 1000,
	}
	expenseRepository := expenseRepository{db: testDb}
	expenseActual, err := expenseRepository.CreateExpense(expenseExpected)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, expenseActual)
	expenseGot, _ := expenseRepository.GetExpenseById(expenseActual.Id)
	assert.Equal(t, 1000, expenseGot.Amount)
}
