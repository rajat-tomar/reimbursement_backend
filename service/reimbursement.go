package service

import (
	"fmt"
	"reimbursement_backend/model"
	"reimbursement_backend/repository"
)

type ReimbursementService interface {
	CreateReimbursement(expense model.Reimbursement) (model.Reimbursement, error)
	GetReimbursements(userId int) ([]model.Reimbursement, error)
	GetUserByEmail(email string) (model.User, error)
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
