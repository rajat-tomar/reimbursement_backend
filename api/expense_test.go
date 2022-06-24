package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

func TestCreateExpense(t *testing.T) {
	t.Run("given expense amount 1000 when valid returns the id of the created expense", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"amount": 1000,
		})

		req, _ := http.NewRequest(http.MethodPost, "/expense", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}
		handler := http.HandlerFunc(expenseController.CreateExpense)
		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		response := model.Response{Data: model.Expense{Id: 7, Amount: 1000}}
		marshalledResponse, _ := json.Marshal(response)
		got := string(body)
		want := string(marshalledResponse)
		assert.JSONEq(t, want, got)
	})
}
