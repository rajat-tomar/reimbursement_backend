package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	requestModel "reimbursement_backend/api/request_model"
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

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var requestBody requestModel.ExpenseRequest
	email := r.Context().Value("email").(string)

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		err = fmt.Errorf("failed to decode request body %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	_, statusCode, err := e.expenseService.CreateExpense(email, requestBody)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var expenses []model.Expense
	email := r.Context().Value("email").(string)
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	category := r.URL.Query().Get("category")

	expenses, statusCode, err := e.expenseService.GetExpenses(email, startDate, endDate, category)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	response.Data = expenses

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}

func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	expenseId, _ := strconv.Atoi(id)
	if expenseId <= 0 {
		response.Message = "either id is empty or invalid"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := e.expenseService.DeleteExpense(expenseId)
	if err != nil {
		response.Message = fmt.Sprintf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "Expense deleted"
	json.NewEncoder(w).Encode(response)
}
