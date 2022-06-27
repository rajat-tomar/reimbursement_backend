package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var expense model.Expense
	var response model.Response
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &expense)

	if expense.Amount == 0 {
		response = model.Response{
			Errors: []error{fmt.Errorf("amount must be greater than 0")},
		}
	} else {
		expense, err := e.expenseService.CreateExpense(expense)
		if err != nil {
			response = model.Response{
				Errors: []error{fmt.Errorf("error creating expense")},
			}
		} else {
			response = model.Response{
				Data: expense,
			}
		}
	}
	json.NewEncoder(w).Encode(response)
}

func (e *expenseController) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var id string
	id = r.URL.Query().Get("id")
	expenseId, _ := strconv.Atoi(id)

	expense, err := e.expenseService.GetExpenseById(expenseId)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}
