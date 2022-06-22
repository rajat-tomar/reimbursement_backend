package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reimbursement_backend/model"
	"testing"
)

func TestGetExpenseById(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/expense", nil)
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
}
