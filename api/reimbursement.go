package api

import (
	"encoding/json"
	"net/http"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
	"strconv"
)

type ReimbursementController interface {
	CreateReimbursement(w http.ResponseWriter, r *http.Request)
	GetReimbursements(w http.ResponseWriter, r *http.Request)
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

func (rmb *reimbursementController) GetReimbursements(w http.ResponseWriter, r *http.Request) {
	var userId int
	id := r.URL.Query().Get("userId")

	if id == "" {
		email := r.Context().Value("email").(string)
		user, err := rmb.reimbursementService.GetUserByEmail(email)
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
	reimbursements, err := rmb.reimbursementService.GetReimbursements(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(reimbursements)
}
