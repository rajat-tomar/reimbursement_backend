package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
	"time"
)

type ExpenseRepository interface {
	CreateExpense(userId int, expense model.Expense) (model.Expense, error)
	GetExpenseById(expenseID int) (model.Expense, error)
	GetExpenses() ([]model.Expense, error)
	DeleteExpense(expenseID int) error
	GetExpensesByDateRange(startDate, endDate time.Time) ([]model.Expense, error)
}

type expenseRepository struct {
	db *sql.DB
}

func (er *expenseRepository) DeleteExpense(expenseID int) error {
	var expense model.Expense

	err := er.db.QueryRow(`SELECT id, amount FROM expenses WHERE id = $1`, expenseID).Scan(&expense.Id, &expense.Amount)
	if err != nil {
		return fmt.Errorf("no expense with given id")
	}
	sqlStatement := `DELETE FROM expenses WHERE Id = $1`
	_, err = er.db.Exec(sqlStatement, expenseID)

	return err
}

func (er *expenseRepository) CreateExpense(userId int, e model.Expense) (model.Expense, error) {
	var expense model.Expense
	sqlStatement := `INSERT INTO expenses(amount, expense_date, category, user_id) VALUES($1, $2, $3, $4) RETURNING amount, expense_date, category`

	err := er.db.QueryRow(sqlStatement, e.Amount, e.ExpenseDate, e.Category, userId).Scan(&expense.Amount, &expense.ExpenseDate, &expense.Category)

	return expense, err
}

func (er *expenseRepository) GetExpenseById(expenseID int) (model.Expense, error) {
	var expense model.Expense

	sqlStatement := `SELECT id, amount FROM expenses WHERE Id = $1`
	err := er.db.QueryRow(sqlStatement, expenseID).Scan(&expense.Id, &expense.Amount)

	return expense, err
}

func (er *expenseRepository) GetExpensesByDateRange(startDate, endDate time.Time) ([]model.Expense, error) {
	var expenses []model.Expense
	sqlStatement := `SELECT id, amount, expense_date, category FROM expenses WHERE expense_date BETWEEN $1 AND $2`

	rows, err := er.db.Query(sqlStatement, startDate, endDate)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(&expense.Id, &expense.Amount, &expense.ExpenseDate, &expense.Category)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (er *expenseRepository) GetExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	sqlStatement := `SELECT id, amount, expense_date, category FROM expenses`

	rows, err := er.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	for rows.Next() {
		var expense model.Expense
		err = rows.Scan(&expense.Id, &expense.Amount, &expense.ExpenseDate, &expense.Category)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func NewExpenseRepository() *expenseRepository {
	return &expenseRepository{
		db: config.GetDb(),
	}
}
