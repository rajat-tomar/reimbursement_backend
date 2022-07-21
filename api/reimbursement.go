package api

import (
	"encoding/json"
	"net/http"
	requestModel "reimbursement_backend/api/request_model"
	"reimbursement_backend/model"
	"reimbursement_backend/service"
	"strconv"
	"time"
)

type ReimbursementController interface {
	CreateReimbursement(w http.ResponseWriter, r *http.Request)
	GetReimbursements(w http.ResponseWriter, r *http.Request)
	ProcessReimbursements(w http.ResponseWriter, r *http.Request)
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
	var reimbursements []model.Reimbursement
	var userId int
	id := r.URL.Query().Get("userId")
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

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
		fetchedReimbursements, err := rmb.reimbursementService.GetReimbursementsByDateRange(userId, startDateTime, endDateTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reimbursements = fetchedReimbursements
	} else {
		fetchedReimbursements, err := rmb.reimbursementService.GetReimbursements(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reimbursements = fetchedReimbursements
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(reimbursements)
}

func (rmb *reimbursementController) ProcessReimbursements(w http.ResponseWriter, r *http.Request) {
	var requestBody requestModel.ProcessReimbursementsRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := rmb.reimbursementService.ProcessReimbursements(requestBody.UserId, requestBody.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
