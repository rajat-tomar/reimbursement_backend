package repository

import (
	"database/sql"
	"reimbursement_backend/model"
)

type ExpenseRepository interface {
	Create(expense model.Expense) (model.Expense, error)
	GetById(expenseID int) (model.Expense, error)
	GetAll() ([]model.Expense, error)
}

type expenseRepository struct {
	db *sql.DB
}

func (er *expenseRepository) Create(e model.Expense) (model.Expense, error) {
	sqlStatement := `INSERT INTO expenses(Amount) VALUES($1)`
	row := er.db.QueryRow(sqlStatement, e.Amount)
	var expense model.Expense
	err := row.Scan(
		&expense.Amount,
	)
	return expense, err
}

func (er *expenseRepository) GetById(expenseID int) (model.Expense, error) {
	sqlStatement := `SELECT id, amount FROM expenses WHERE Id = $1`
	row := er.db.QueryRow(sqlStatement, expenseID)
	var expense model.Expense
	err := row.Scan(
		&expense.Id,
		&expense.Amount,
	)
	return expense, err
}
