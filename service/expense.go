package service

import (
	"fmt"
	"net/http"
	"reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
	"strconv"
	"time"
)

type ExpenseService interface {
	GetExpenseById(expenseId int) (model.Expense, int, error)
	CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, int, error)
	GetExpenses(userId int, category string) ([]model.Expense, error)
	GetExpensesByDateRange(userId int, startDate, endDate time.Time, category string) ([]model.Expense, error)
	DeleteExpense(email string, expenseId int) (int, error)
	UpdateExpense(expenseId string, requestBody request_model.ExpenseRequest) (int, error)
	GetUserByEmail(email string) (model.User, error)
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
	userService       UserService
}

func NewExpenseService() *expenseService {
	return &expenseService{
		expenseRepository: repository.NewExpenseRepository(),
		userService:       NewUserService(),
	}
}

func (es *expenseService) GetExpenseById(expenseId int) (model.Expense, int, error) {
	if expenseId <= 0 {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("expense id must be greater than 0")
	}
	expense, err := es.expenseRepository.GetExpenseById(expenseId)
	if err != nil {
		return model.Expense{}, http.StatusNotFound, fmt.Errorf("failed to get expense: %v", err)
	}

	return expense, http.StatusOK, nil
}

func (es *expenseService) CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, int, error) {
	var expense model.Expense
	user, err := es.userService.GetUserByEmail(email)
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

func (es *expenseService) GetExpenses(userId int, category string) ([]model.Expense, error) {
	if userId <= 0 {
		return nil, fmt.Errorf("user id must be greater than 0")
	}
	expenses, err := es.expenseRepository.GetExpenses(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses: %v", err)
	}
	if category != "" {
		var filteredExpenses []model.Expense
		for _, expense := range expenses {
			if expense.Category == category {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}

		expenses = filteredExpenses
	} else {
		var filteredExpenses []model.Expense
		for _, expense := range expenses {
			if expense.Status == "pending" {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}

		expenses = filteredExpenses
	}

	return expenses, nil
}

func (es *expenseService) GetExpensesByDateRange(userId int, startDate, endDate time.Time, category string) ([]model.Expense, error) {
	if userId <= 0 {
		return []model.Expense{}, fmt.Errorf("user id must be greater than 0")
	}
	expenses, err := es.expenseRepository.GetExpensesByDateRange(userId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses: %v", err)
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

	return expenses, nil
}

func (es *expenseService) DeleteExpense(email string, expenseId int) (int, error) {
	user, err := es.userService.GetUserByEmail(email)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("no user found with email %s: %v", email, err)
	}
	userId := user.Id

	if expenseId <= 0 {
		return http.StatusBadRequest, fmt.Errorf("expense id must be greater than 0")
	}
	err = es.expenseRepository.DeleteExpense(userId, expenseId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete expense: %v", err)
	}

	return http.StatusNoContent, nil
}

func (es *expenseService) UpdateExpense(expenseId string, requestBody request_model.ExpenseRequest) (int, error) {
	var expense model.Expense

	if expenseId == "" {
		return http.StatusBadRequest, fmt.Errorf("expense id is required")
	}
	expenseIdInt, err := strconv.Atoi(expenseId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to convert expense id to int: %v", err)
	}
	expense.Id = expenseIdInt
	expense.Status = requestBody.Status
	err = es.expenseRepository.UpdateExpense(expense)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update expense: %v", err)
	}

	return http.StatusNoContent, nil
}

func (es *expenseService) GetUserByEmail(email string) (model.User, error) {
	user, err := es.userService.GetUserByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}
