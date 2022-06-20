package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/model"
	"testing"
)

type mockExpenseRepository struct {
}

func (m mockExpenseRepository) GetById(expenseID int) (model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) GetAll() ([]model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseRepository) Create(expense model.Expense) (model.Expense, error) {
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
	expenseService := expenseService{
		expenseRepository: &mockExpenseRepository{},
	}
	_, err := expenseService.Create(expense)
	assert.Equal(t, errors.New("error creating expense"), err)
}

func TestCreate_ChecksForSuccessfulCreation(t *testing.T) {
	expenseExpected := model.Expense{
		Id:     0,
		Amount: 200,
	}
	expenseService := expenseService{
		expenseRepository: &mockExpenseRepository{},
	}
	expenseActual, err := expenseService.Create(expenseExpected)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, expenseActual)
	assert.Equal(t, expenseExpected, expenseActual)
}
