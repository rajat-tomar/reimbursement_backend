package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type ExpenseService interface {
	CreateExpense(expense model.Expense) (model.Expense, error)
	GetExpenses() ([]model.Expense, error)
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

func (es *expenseService) GetExpenses() ([]model.Expense, error) {
	expenses, err := es.expenseRepository.GetExpenses()
	if err != nil {
		return nil, fmt.Errorf("error from repo: %w", err)
	}
	return expenses, nil
}

func NewExpenseService() *expenseService {
	return &expenseService{
		expenseRepository: repository.NewExpenseRepository(),
	}
}
