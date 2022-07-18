package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"testing"
	"time"
)

type mockExpenseService struct{}

func (m mockExpenseService) GetExpenseById(expenseId int) (model.Expense, int, error) {
	if expenseId <= 0 {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("expense id must be greater than 0")
	}
	if expenseId == 2 {
		return model.Expense{}, http.StatusNotFound, nil
	}
	return model.Expense{Id: 1, Amount: 1000, Category: "Learning and Development", UserId: 1, Status: "pending"}, http.StatusOK, nil
}

func (m mockExpenseService) CreateExpense(email string, requestBody request_model.ExpenseRequest) (model.Expense, int, error) {
	if requestBody.Amount <= 0 {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("amount must be greater than zero")
	}
	if requestBody.ExpenseDate == "" {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("expense date can't be empty")
	}
	if requestBody.Category == "" {
		return model.Expense{}, http.StatusBadRequest, fmt.Errorf("category can't be empty")
	}
	expenseDate, _ := time.Parse("2006-01-02", requestBody.ExpenseDate)
	return model.Expense{Id: 1, Amount: 1000, Category: "Learning and Development", ExpenseDate: expenseDate, UserId: 1, Status: "pending"}, http.StatusCreated, nil
}

func (m mockExpenseService) GetExpenses(email, startDate, endDate, category string, userId int) ([]model.Expense, int, error) {
	if userId == 3 {
		return nil, http.StatusNotFound, fmt.Errorf("user not found")
	}
	return []model.Expense{
		{Id: 1, Amount: 1000, Category: "Learning and Development", UserId: 1, Status: "pending"},
		{Id: 2, Amount: 2000, Category: "Learning and Development", UserId: 1, Status: "pending"},
	}, http.StatusOK, nil
}

func (m mockExpenseService) DeleteExpense(email string, expenseId int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockExpenseService) UpdateExpense(expenseId string, requestBody request_model.ExpenseRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func TestGetExpenseById(t *testing.T) {
	t.Run("should return expense when found for given id", func(t *testing.T) {
		var expense model.Expense
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=1", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.GetExpenseById)

		handler.ServeHTTP(rr, req)
		err := json.NewDecoder(rr.Body).Decode(&expense)
		got := expense
		want := model.Expense{Id: 1, Amount: 1000, Category: "Learning and Development", UserId: 1, Status: "pending"}
		status := rr.Code

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, want, got)
	})

	t.Run("should return status 400 when invalid id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=0", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.GetExpenseById)

		handler.ServeHTTP(rr, req)
		status := rr.Code

		assert.Equal(t, http.StatusBadRequest, status)
	})

	t.Run("should return status 404 when expense not found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=2", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.GetExpenseById)

		handler.ServeHTTP(rr, req)
		status := rr.Code

		assert.Equal(t, http.StatusNotFound, status)
	})
}

func TestCreateExpense(t *testing.T) {
	t.Run("should return status 500 when invalid request body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer([]byte(`{"amount":0,category:"","expense_date":"","user_id":0}`)))
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.CreateExpense)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should return status 201 when expense successfully created", func(t *testing.T) {
		reqBody := []byte(`{"amount":1000,"category":"Learning and Development","expense_date":"2020-01-01"}`)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer(reqBody))
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.CreateExpense)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("should return status 400 when amount is not provided or invalid", func(t *testing.T) {
		reqBody := []byte(`{"amount":0,"category":"Learning and Development","expense_date":"2020-01-01"}`)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer(reqBody))
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.CreateExpense)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return status 400 when category is not provided or invalid", func(t *testing.T) {
		reqBody := []byte(`{"amount":1000,"expense_date":"2020-01-01","category":""}`)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer(reqBody))
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.CreateExpense)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return status 400 when expense date is not provided or invalid", func(t *testing.T) {
		reqBody := []byte(`{"amount":1000,"category":"Learning and Development"}`)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer(reqBody))
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.CreateExpense)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestGetExpenses(t *testing.T) {
	t.Run("should return status 200 when expenses successfully fetched", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/expenses?userId=1", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.GetExpenses)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return status 404 when no expenses found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/expenses?userId=3", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{expenseService: &mockExpenseService{}}
		handler := http.HandlerFunc(expenseController.GetExpenses)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "email", "rajat.rajattomar@gmail.com")
		handler.ServeHTTP(rr, req.WithContext(ctx))

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
