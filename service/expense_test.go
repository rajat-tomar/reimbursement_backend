package service

import (
	"errors"
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

func (m mockExpenseRepository) Create(e model.Expense) (int, error) {
	if e.Amount == -1 {
		return 0, errors.New("error creating expense")
	}

	return 7, nil
}

func Test_Expense_Service_Create_Returns_Error_When_Unable_To_Create(t *testing.T) {
	expense := model.Expense{
		ID:     0,
		Amount: -1,
	}

	expenseService := expenseService{
		expenseRepository: &mockExpenseRepository{},
	}

	_, err := expenseService.Create(expense)

	assert.Equal(t, errors.New("error creating expense"), err)
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
