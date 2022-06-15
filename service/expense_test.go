package service

import (
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/model"
	"testing"
)

type mockExpenseRepository struct {
}

func (m mockExpenseRepository) GetById(expenseID int) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) GetAll() ([]model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) Create(expense model.Expense) (int, error) {
	return 7, nil
}

func TestCreateExpense(t *testing.T) {
	expense := model.Expense{
		ID:     0,
		Amount: 200.0,
	}

	expenseService := expenseService{
		expenseRepository: &mockExpenseRepository{},
	}
	expenseId, err := expenseService.Create(expense)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, expenseId)
	assert.Equal(t, 7, expenseId)
}
