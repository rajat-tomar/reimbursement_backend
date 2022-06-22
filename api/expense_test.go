package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reimbursement_backend/model"
	"testing"
)

func TestGetExpenseById(t *testing.T) {
	t.Run("returns expense with id 20", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/expense?id=20", nil)
		response := httptest.NewRecorder()
		GetExpenseById(response, request)
		var expense model.Expense
		err := json.NewDecoder(response.Body).Decode(&expense)
		got := expense
		want := model.Expense{
			Id:     20,
			Amount: 1000,
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, want, got)
	})

	t.Run("returns expense with id 10", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/expense?id=10", nil)
		response := httptest.NewRecorder()
		GetExpenseById(response, request)
		var expense model.Expense
		err := json.NewDecoder(response.Body).Decode(&expense)
		got := expense
		want := model.Expense{
			Id:     10,
			Amount: 500,
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, want, got)
	})
}
