package api

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
)

type ReimbursementController interface {
	CreateReimbursement(w http.ResponseWriter, r *http.Request)
}

type reimbursementController struct {
	reimbursementService service.ReimbursementService
}

func NewReimbursementController() *reimbursementController {
	return &reimbursementController{
		reimbursementService: service.NewReimbursementService(),
	}
}

func (rmb *reimbursementController) CreateReimbursement(w http.ResponseWriter, r *http.Request) {
	var reimbursement model.Reimbursement
	if err := json.NewDecoder(r.Body).Decode(&reimbursement); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdReimbursement, err := rmb.reimbursementService.CreateReimbursement(reimbursement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(createdReimbursement)
}
