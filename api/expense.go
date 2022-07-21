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
	GetExpenseById(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	GetExpenses(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
}

type expenseController struct {
	expenseService service.ExpenseService
}

func NewExpenseController() *expenseController {
	return &expenseController{
		expenseService: service.NewExpenseService(),
	}
}

func (e *expenseController) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	id := r.URL.Query().Get("id")
	expenseId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expense, statusCode, err := e.expenseService.GetExpenseById(expenseId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(expense)
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var requestBody requestModel.ExpenseRequest
	email := r.Context().Value("email").(string)

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		err = fmt.Errorf("failed to decode request body %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, statusCode, err := e.expenseService.CreateExpense(email, requestBody)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	var expenses []model.Expense
	var userId int
	id := r.URL.Query().Get("userId")
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	category := r.URL.Query().Get("category")

	if id == "" {
		email := r.Context().Value("email").(string)
		user, err := e.expenseService.GetUserByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userId = user.Id
	} else {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userId = idInt
	}
	if startDate != "" && endDate != "" {
		startDateTime, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		endDateTime, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fetchedExpenses, err := e.expenseService.GetExpensesByDateRange(userId, startDateTime, endDateTime, category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		expenses = fetchedExpenses
	} else {
		fetchedExpenses, err := e.expenseService.GetExpenses(userId, category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		expenses = fetchedExpenses
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(expenses)
}

func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	id := r.URL.Query().Get("id")
	expenseId, _ := strconv.Atoi(id)

	statusCode, err := e.expenseService.DeleteExpense(email, expenseId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
}

func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var requestBody requestModel.ExpenseRequest
	expenseId := r.URL.Query().Get("id")

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		err = fmt.Errorf("failed to decode request body %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	statusCode, err := e.expenseService.UpdateExpense(expenseId, requestBody)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
}
