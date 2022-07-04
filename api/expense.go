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
	GetExpenses(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
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
		response = model.Response{Message: "Content-Type must be application/json"}
		json.NewEncoder(w).Encode(response)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&expense)
	if err != nil {
		response.Message = fmt.Sprintf("error from json: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if expense.Amount <= 0 {
		response.Message = "Amount must be greater than 0"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	expense, err = e.expenseService.CreateExpense(expense)
	if err != nil {
		response.Message = fmt.Sprintf("error from service: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = expense
	json.NewEncoder(w).Encode(response)
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	w.Header().Set("Content-Type", "application/json")

	expenses, err := e.expenseService.GetExpenses()
	if err != nil {
		response.Message = fmt.Sprintf("error from service: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.Data)
		return
	}

	response.Data = expenses
	json.NewEncoder(w).Encode(response)
}

func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteExpense")
	var response model.Response
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	amount := r.URL.Query().Get("amount")
	amountInt, _ := strconv.Atoi(amount)
	fmt.Printf("amount: %d", amountInt)
	expenseId, _ := strconv.Atoi(id)
	fmt.Printf("expenseID: %d", expenseId)
	if expenseId <= 0 {
		response.Message = "id is required"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := e.expenseService.DeleteExpense(expenseId)
	if err != nil {
		response.Message = fmt.Sprintf("error from service: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "Expense deleted"
	json.NewEncoder(w).Encode(response)
}

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}
