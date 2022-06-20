package service

import (
	"errors"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type ExpenseService interface {
	Create(expense model.Expense) (model.Expense, error)
	GetById(expenseId int) (model.Expense, error)
	GetAll() ([]model.Expense, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func (es *expenseService) Create(ex model.Expense) (model.Expense, error) {
	expense, err := es.expenseRepository.Create(ex)
	if err != nil {
		return model.Expense{}, errors.New("error creating expense")
	}
	return expense, nil
}
