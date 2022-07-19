package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type ReimbursementRepository interface {
	CreateReimbursement(reimbursement model.Reimbursement) (model.Reimbursement, error)
}

type reimbursementRepository struct {
	db *sql.DB
}

func NewReimbursementRepository() *reimbursementRepository {
	return &reimbursementRepository{
		db: config.GetDb(),
	}
}

func (rmb *reimbursementRepository) CreateReimbursement(reimbursement model.Reimbursement) (model.Reimbursement, error) {
	var createdReimbursement model.Reimbursement
	sqlStatement := `INSERT INTO reimbursements (amount, user_id, expense_id, status) VALUES($1, $2, $3, $4) RETURNING id, amount, user_id, expense_id, status, processed_on`

	err := rmb.db.QueryRow(sqlStatement, reimbursement.Amount, reimbursement.UserId, reimbursement.ExpenseId, "pending").Scan(&createdReimbursement.Id, &createdReimbursement.Amount, &createdReimbursement.UserId, &createdReimbursement.ExpenseId, &createdReimbursement.Status, &createdReimbursement.ProcessedOn)
	if err != nil {
		return model.Reimbursement{}, fmt.Errorf("error creating reimbursement: %v", err)
	}

	return createdReimbursement, nil
}
