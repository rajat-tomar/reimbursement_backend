package repository

import (
	"database/sql"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type ExpenseRepository interface {
	CreateExpense(expense model.Expense) (model.Expense, error)
	GetExpenseById(expenseID int) (model.Expense, error)
}

type expenseRepository struct {
	db *sql.DB
}

func (er *expenseRepository) CreateExpense(e model.Expense) (model.Expense, error) {
	sqlStatement := `INSERT INTO expenses(amount) VALUES($1) RETURNING Id, Amount`
	var expense model.Expense
	err := er.db.QueryRow(sqlStatement, e.Amount).Scan(&expense.Id, &expense.Amount)
	return expense, err
}

func (er *expenseRepository) GetExpenseById(expenseID int) (model.Expense, error) {
	sqlStatement := `SELECT id, amount FROM expenses WHERE Id = $1`
	var expense model.Expense
	err := er.db.QueryRow(sqlStatement, expenseID).Scan(&expense.Id, &expense.Amount)
	return expense, err
}

func NewExpenseRepository() *expenseRepository {
	return &expenseRepository{
		db: config.GetDb(),
	}
}
