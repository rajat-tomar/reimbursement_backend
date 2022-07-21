package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
	"time"
)

type ReimbursementService interface {
	CreateReimbursement(expense model.Reimbursement) (model.Reimbursement, error)
	GetReimbursements(userId int) ([]model.Reimbursement, error)
	GetUserByEmail(email string) (model.User, error)
	ProcessReimbursements(userId int, status string) error
	GetReimbursementsByDateRange(userId int, startDate, endDate time.Time) ([]model.Reimbursement, error)
}

type reimbursementService struct {
	userService             UserService
	reimbursementRepository repository.ReimbursementRepository
}

func NewReimbursementService() *reimbursementService {
	return &reimbursementService{
		userService:             NewUserService(),
		reimbursementRepository: repository.NewReimbursementRepository(),
	}
}

func (rmb *reimbursementService) CreateReimbursement(reimbursement model.Reimbursement) (model.Reimbursement, error) {
	createdReimbursement, err := rmb.reimbursementRepository.CreateReimbursement(reimbursement)
	if err != nil {
		return model.Reimbursement{}, fmt.Errorf("error creating reimbursement: %v", err)
	}

	return createdReimbursement, nil
}

func (rmb *reimbursementService) GetReimbursements(userId int) ([]model.Reimbursement, error) {
	if userId <= 0 {
		return nil, fmt.Errorf("user id must be greater than 0")
	}
	reimbursements, err := rmb.reimbursementRepository.GetReimbursements(userId)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %v", err)
	}

	var filteredReimbursements []model.Reimbursement
	for _, reimbursement := range reimbursements {
		if reimbursement.Status == "pending" {
			filteredReimbursements = append(filteredReimbursements, reimbursement)
		}
	}

	return filteredReimbursements, nil
}

func (rmb *reimbursementService) GetUserByEmail(email string) (model.User, error) {
	user, err := rmb.userService.GetUserByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (rmb *reimbursementService) ProcessReimbursements(userId int, status string) error {
	err := rmb.reimbursementRepository.ProcessReimbursements(userId, status)
	if err != nil {
		return fmt.Errorf("error processing reimbursements: %v", err)
	}

	return nil
}

func (rmb *reimbursementService) GetReimbursementsByDateRange(userId int, startDate, endDate time.Time) ([]model.Reimbursement, error) {
	reimbursements, err := rmb.reimbursementRepository.GetReimbursementsByDateRange(userId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error getting reimbursements: %v", err)
	}

	return reimbursements, nil
}
