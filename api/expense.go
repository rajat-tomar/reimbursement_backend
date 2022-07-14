package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	requestModel "reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
	"strconv"
	"time"
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
	if requestBody.Amount <= 0 {
		err := "Amount must be greater than 0"
		http.Error(w, err, http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	if requestBody.Category == "" {
		err := "Category can't be empty"
		http.Error(w, err, http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	_, err := e.expenseService.CreateExpense(email, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var expenses []model.Expense
	w.Header().Set("Content-Type", "application/json")
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	category := r.URL.Query().Get("category")

	if startDate != "" && endDate != "" {
		startDateTime, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			response.Message = fmt.Sprintf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		endDateTime, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			response.Message = fmt.Sprintf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		fetchedExpenses, err := e.expenseService.GetExpensesByDateRange(startDateTime, endDateTime)
		if err != nil {
			response.Message = fmt.Sprintf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}
		expenses = fetchedExpenses
	} else {
		fetchedExpenses, err := e.expenseService.GetExpenses()
		if err != nil {
			response.Message = fmt.Sprintf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}
		expenses = fetchedExpenses
	}

	if category != "" && expenses != nil {
		var filteredExpenses []model.Expense
		for _, expense := range expenses {
			if expense.Category == category {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}
		response.Data = filteredExpenses
	} else {
		response.Data = expenses
	}
	json.NewEncoder(w).Encode(response)
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
