package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type ReimbursementRepository interface {
	CreateReimbursement(reimbursement model.Reimbursement) (model.Reimbursement, error)
	GetReimbursements(userId int) ([]model.Reimbursement, error)
	ProcessReimbursements(userId int, status string) error
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
	sqlStatement := `INSERT INTO reimbursements (amount, user_id, expense_id, category, status) VALUES($1, $2, $3, $4, $5) RETURNING id, amount, user_id, expense_id, category, status`

	err := rmb.db.QueryRow(sqlStatement, reimbursement.Amount, reimbursement.UserId, reimbursement.ExpenseId, reimbursement.Category, "pending").Scan(&createdReimbursement.Id, &createdReimbursement.Amount, &createdReimbursement.UserId, &createdReimbursement.ExpenseId, &reimbursement.Category, &createdReimbursement.Status)
	if err != nil {
		return model.Reimbursement{}, fmt.Errorf("error creating reimbursement: %v", err)
	}

	return createdReimbursement, nil
}

func (rmb *reimbursementRepository) GetReimbursements(userId int) ([]model.Reimbursement, error) {
	var reimbursements []model.Reimbursement
	sqlStatement := `SELECT id, amount, user_id, expense_id, category, status FROM reimbursements WHERE user_id = $1`

	rows, err := rmb.db.Query(sqlStatement, userId)
	if err != nil {
		return nil, fmt.Errorf("repo: error getting reimbursements: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			config.Logger.Panicf("repo: error closing rows: %v", err)
		}
	}(rows)
	for rows.Next() {
		var reimbursement model.Reimbursement

		err := rows.Scan(&reimbursement.Id, &reimbursement.Amount, &reimbursement.UserId, &reimbursement.ExpenseId, &reimbursement.Category, &reimbursement.Status)
		if err != nil {
			return nil, fmt.Errorf("repo: error scanning reimbursements: %v", err)
		}

		reimbursements = append(reimbursements, reimbursement)
	}

	return reimbursements, nil
}

func (rmb *reimbursementRepository) ProcessReimbursements(userId int, status string) error {
	sqlStatement := `UPDATE reimbursements SET status = $1 WHERE user_id = $2 AND status = $3`

	_, err := rmb.db.Exec(sqlStatement, status, userId, "pending")
	if err != nil {
		return fmt.Errorf("repo: error processing reimbursements: %v", err)
	}

	return nil
}
