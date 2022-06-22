package handler

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
)

func GetExpenseById(w http.ResponseWriter, r *http.Request) {
	expense := model.Expense{Id: 20, Amount: 1000}
	json.NewEncoder(w).Encode(expense)
}
