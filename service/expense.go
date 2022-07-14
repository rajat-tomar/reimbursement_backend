package service

import (
	"fmt"
	"net/http"
	"reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
	"time"
)

type ExpenseService interface {
	CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, int, error)
	GetExpenses(email, startDate, endDate, category string) ([]model.Expense, int, error)
	DeleteExpense(id int) error
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
	userService       UserService
}

func (es *expenseService) CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, int, error) {
	var expense model.Expense
	user, err := es.userService.FindByEmail(email)
	if err != nil {
		return model.Expense{}, http.StatusInternalServerError, fmt.Errorf("failed to create expense: %v", err)
	}

	if requestBody.Amount <= 0 {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("amount must be greater than 0")
	}
	if requestBody.Category == "" {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("category can't be empty")
	}
	userId := user.Id
	expenseDate, _ := time.Parse("2006-01-02", requestBody.ExpenseDate)
	expense.Amount = requestBody.Amount
	expense.Category = requestBody.Category
	expense.ExpenseDate = expenseDate
	expense.UserId = userId
	createdExpense, err := es.expenseRepository.CreateExpense(userId, expense)
	if err != nil {
		return model.Expense{}, http.StatusFailedDependency, fmt.Errorf("failed to create expense: %v", err)
	}

	return createdExpense, http.StatusCreated, nil
}

func (es *expenseService) GetExpenses(email, startDate, endDate, category string) ([]model.Expense, int, error) {
	var expenses []model.Expense
	user, err := es.userService.FindByEmail(email)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("no user found with email %s: %v", email, err)
	}
	userId := user.Id

	if startDate != "" && endDate != "" {
		startDateTime, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, http.StatusBadRequest, fmt.Errorf("failed to parse start date: %v", err)
		}
		endDateTime, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, http.StatusBadRequest, fmt.Errorf("failed to parse end date: %v", err)
		}
		fetchedExpenses, err := es.expenseRepository.GetExpensesByDateRange(userId, startDateTime, endDateTime)
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("failed to get expenses: %v", err)
		}

		expenses = fetchedExpenses
	} else {
		fetchedExpenses, err := es.expenseRepository.GetExpenses(userId)
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("failed to get expenses: %v", err)
		}

		expenses = fetchedExpenses
	}
	if category != "" {
		var filteredExpenses []model.Expense
		for _, expense := range expenses {
			if expense.Category == category {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}

		expenses = filteredExpenses
	}

	return expenses, http.StatusOK, nil
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
