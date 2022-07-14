package service

import (
	"fmt"
	"reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
	"time"
)

type ExpenseService interface {
	CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, error)
	GetExpenses() ([]model.Expense, error)
	DeleteExpense(id int) error
	GetExpensesByDateRange(startDate, endDate time.Time) ([]model.Expense, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
	userService       UserService
}

func (es *expenseService) CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, error) {
	var expense model.Expense
	user, err := es.userService.FindByEmail(email)
	if err != nil {
		return model.Expense{}, fmt.Errorf("failed to create expense: %v", err)
	}

	userId := user.Id
	expenseDate, _ := time.Parse("2006-01-02", requestBody.ExpenseDate)
	expense.Amount = requestBody.Amount
	expense.Category = requestBody.Category
	expense.ExpenseDate = expenseDate
	expense.UserId = userId
	createdExpense, err := es.expenseRepository.CreateExpense(userId, expense)
	if err != nil {
		return model.Expense{}, fmt.Errorf("failed to create expense: %v", err)
	}

	return createdExpense, nil
}

func (es *expenseService) GetExpenses() ([]model.Expense, error) {
	expenses, err := es.expenseRepository.GetExpenses()
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) GetExpensesByDateRange(startDate, endDate time.Time) ([]model.Expense, error) {
	expenses, err := es.expenseRepository.GetExpensesByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *expenseService) DeleteExpense(id int) error {
	err := es.expenseRepository.DeleteExpense(id)
	if err != nil {
		return fmt.Errorf("failed to delete expense: %v", err)
	}

	return nil
}

func NewExpenseService() *expenseService {
	return &expenseService{
		expenseRepository: repository.NewExpenseRepository(),
		userService:       NewUserService(),
	}
}
