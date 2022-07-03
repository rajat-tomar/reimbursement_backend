package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type ExpenseService interface {
	CreateExpense(expense model.Expense) (model.Expense, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func (es *expenseService) CreateExpense(e model.Expense) (model.Expense, error) {
	expense, err := es.expenseRepository.CreateExpense(e)
	if err != nil {
		return model.Expense{}, fmt.Errorf("error from repo: %w", err)
	}
	return expense, nil
}

func NewExpenseService() *expenseService {
	return &expenseService{
		expenseRepository: repository.NewExpenseRepository(),
	}
}
