package service

import (
	"errors"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

const errorMessageForCreateExpense = "error creating expense"

type ExpenseService interface {
	Create(expense model.Expense) (int, error)
	GetById(expenseId int) (model.Expense, error)
	GetAll() ([]model.Expense, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func (es *expenseService) Create(expense model.Expense) (int, error) {
	expenseId, err := es.expenseRepository.Create(expense)
	if err != nil {
		return 0, errors.New(errorMessageForCreateExpense)
	}
	return expenseId, nil
}
