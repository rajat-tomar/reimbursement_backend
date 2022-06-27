package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type ExpenseService interface {
	CreateExpense(expense model.Expense) (model.Expense, error)
	GetExpenseById(expenseId int) (model.Expense, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func (es *expenseService) GetExpenseById(expenseId int) (model.Expense, error) {
	return es.expenseRepository.GetExpenseById(expenseId)
}

func (es *expenseService) CreateExpense(e model.Expense) (model.Expense, error) {
	expense, err := es.expenseRepository.CreateExpense(e)
	if err != nil {
		return model.Expense{}, fmt.Errorf("error creating expense")
	}
	return expense, nil
}

func NewExpenseService() *expenseService {
	return &expenseService{
		expenseRepository: repository.NewExpenseRepository(),
	}
}
