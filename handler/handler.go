package handler

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"strconv"
)

func GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	q := r.URL.Query()
	id, _ := strconv.Atoi(q.Get("id"))
	if id == 20 {
		expense = model.Expense{Id: 20, Amount: 1000}
	}
	if id == 10 {
		expense = model.Expense{Id: 10, Amount: 500}
	}
	json.NewEncoder(w).Encode(expense)
}
