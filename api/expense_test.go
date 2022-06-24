package api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reimbursement_backend/model"
	"testing"
)

type mockExpenseService struct{}

func (e *mockExpenseService) Create(expense model.Expense) (model.Expense, error) {
	//TODO implement me
	panic("implement me")
}

func (e *mockExpenseService) GetExpenseById(expenseId int) (model.Expense, error) {
	if expenseId == 999 {
		return model.Expense{}, fmt.Errorf("expense not found")
	}

	if expenseId == 2 {
		return model.Expense{Id: 2, Amount: 2000}, nil
	}

	return model.Expense{Id: 1, Amount: 1000}, nil
}

func TestGetExpenseById(t *testing.T) {
	t.Run("returns expense for id 1", func(t *testing.T) {
		var expense model.Expense
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=1", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}

		handler := http.HandlerFunc(expenseController.GetExpenseById)
		handler.ServeHTTP(rr, req)
		err := json.NewDecoder(rr.Body).Decode(&expense)
		got := expense
		want := model.Expense{
			Id:     1,
			Amount: 1000,
		}
		status := rr.Code

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, want, got)
	})

	t.Run("returns expense for id 2", func(t *testing.T) {
		var expense model.Expense
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=2", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}

		handler := http.HandlerFunc(expenseController.GetExpenseById)
		handler.ServeHTTP(rr, req)
		err := json.NewDecoder(rr.Body).Decode(&expense)
		got := expense
		want := model.Expense{
			Id:     2,
			Amount: 2000,
		}
		status := rr.Code

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, want, got)
	})

	t.Run("for a given id when expense not found returns 404", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/expense?id=999", nil)
		rr := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}

		handler := http.HandlerFunc(expenseController.GetExpenseById)
		handler.ServeHTTP(rr, req)
		want := http.StatusNotFound
		got := rr.Code

		assert.Equal(t, want, got)
	})
}
