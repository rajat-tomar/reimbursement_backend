package api

import (
	"encoding/json"
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
	return model.Expense{
		Id:     1,
		Amount: 1000,
	}, nil
}

func TestGetExpenseById(t *testing.T) {
	t.Run("returns expense with id 10", func(t *testing.T) {
		var expense model.Expense
		request, _ := http.NewRequest(http.MethodGet, "/expense?id=1", nil)
		response := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}
		expenseController.GetExpenseById(response, request)

		err := json.NewDecoder(response.Body).Decode(&expense)
		got := expense
		want := model.Expense{
			Id:     1,
			Amount: 1000,
		}

		assert.Equal(t, nil, err)
		assert.Equal(t, want, got)
	})

	t.Run("return status ok on success request", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/expense?id=2", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		expenseController := expenseController{
			expenseService: &mockExpenseService{},
		}

		handler := http.HandlerFunc(expenseController.GetExpenseById)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		assert.Equal(t, http.StatusOK, status)
	})
}
