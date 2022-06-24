package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
	"strconv"
)

type ExpenseController interface {
	CreateExpense(w http.ResponseWriter, r *http.Request)
	GetExpenseById(w http.ResponseWriter, r *http.Request)
}

type expenseController struct {
	expenseService service.ExpenseService
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var reqExpense model.Expense
	json.NewDecoder(r.Body).Decode(&reqExpense)
	resExpense, err := e.expenseService.Create(reqExpense)
	var response string
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response = fmt.Sprintf("Expense amount %d is created successfully with %d", resExpense.Amount, resExpense.Id)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		response = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}

func (e *expenseController) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var id string
	id = r.URL.Query().Get("id")
	expenseId, _ := strconv.Atoi(id)

	result, err := e.expenseService.GetExpenseById(expenseId)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}
