package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/model"
	"testing"
)

type mockExpenseRepository struct {
}

func (m mockExpenseRepository) GetExpenseById(expenseID int) (model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) GetAll() ([]model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	if expense.Amount == -1 {
		return model.Expense{}, errors.New("error creating expense")
	}
	return expense, nil
}

func TestCreate_ReturnsErrorWhenUnableToCreate(t *testing.T) {
	expense := model.Expense{
		Id:     0,
		Amount: -1,
	}
	expenseService := expenseService{expenseRepository: &mockExpenseRepository{}}
	_, err := expenseService.CreateExpense(expense)
	assert.EqualError(t, err, "error from repo: error creating expense")
}

func TestCreate_ChecksForSuccessfulCreation(t *testing.T) {
	expenseExpected := model.Expense{
		Id:     0,
		Amount: 200,
	}
	expenseService := expenseService{expenseRepository: &mockExpenseRepository{}}
	expenseActual, err := expenseService.CreateExpense(expenseExpected)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, expenseActual)
	assert.Equal(t, expenseExpected, expenseActual)
}
