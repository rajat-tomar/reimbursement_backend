package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
)

type ExpenseController interface {
	CreateExpense(w http.ResponseWriter, r *http.Request)
}

type expenseController struct {
	expenseService service.ExpenseService
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	var response model.Response
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		response = model.Response{Status: "error", Errors: []error{fmt.Errorf("Content-Type must be application/json")}}
		json.NewEncoder(w).Encode(response)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&expense)
	if err != nil {
		response.Status = "error"
		response.Errors = append(response.Errors, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if expense.Amount <= 0 {
		response.Status = "error"
		response.Errors = append(response.Errors, fmt.Errorf("amount must be greater than 0"))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	expense, err = e.expenseService.CreateExpense(expense)
	if err != nil {
		response.Status = "error"
		w.WriteHeader(http.StatusInternalServerError)
		response.Errors = append(response.Errors, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = "success"
	response.Data = expense
	json.NewEncoder(w).Encode(response)
}

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}
