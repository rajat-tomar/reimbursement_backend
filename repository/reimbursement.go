package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
	"time"
)

type ReimbursementRepository interface {
	CreateReimbursement(reimbursement model.Reimbursement) (model.Reimbursement, error)
	GetReimbursements(userId int) ([]model.Reimbursement, error)
	ProcessReimbursements(userId int, status string) error
	GetReimbursementsByDateRange(userId int, startDate, endDate time.Time) ([]model.Reimbursement, error)
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
	sqlStatement := `UPDATE reimbursements SET status = $1, processed_on = $2 WHERE user_id = $3 AND status = $4`

	_, err := rmb.db.Exec(sqlStatement, status, time.Now(), userId, "pending")
	if err != nil {
		return fmt.Errorf("repo: error processing reimbursements: %v", err)
	}

	return nil
}

func (rmb *reimbursementRepository) GetReimbursementsByDateRange(userId int, startDate, endDate time.Time) ([]model.Reimbursement, error) {
	var reimbursements []model.Reimbursement
	sqlStatement := `SELECT id, amount, user_id, expense_id, category, status, processed_on FROM reimbursements WHERE user_id = $1 AND processed_on BETWEEN $2 AND $3`

	rows, err := rmb.db.Query(sqlStatement, userId, startDate, endDate)
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

		err := rows.Scan(&reimbursement.Id, &reimbursement.Amount, &reimbursement.UserId, &reimbursement.ExpenseId, &reimbursement.Category, &reimbursement.Status, &reimbursement.ProcessedOn)
		if err != nil {
			return nil, fmt.Errorf("repo: error scanning reimbursements: %v", err)
		}

		reimbursements = append(reimbursements, reimbursement)
	}

	return reimbursements, nil
}
