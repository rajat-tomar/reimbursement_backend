package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
	"time"
)

type ExpenseRepository interface {
	GetExpenseById(expenseID int) (model.Expense, error)
	CreateExpense(userId int, expense model.Expense) (model.Expense, error)
	GetExpenses(userId int) ([]model.Expense, error)
	GetExpensesByDateRange(userId int, startDate, endDate time.Time) ([]model.Expense, error)
	DeleteExpense(userId, expenseId int) error
}

type expenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository() *expenseRepository {
	return &expenseRepository{
		db: config.GetDb(),
	}
}

func (er *expenseRepository) GetExpenseById(expenseID int) (model.Expense, error) {
	var expense model.Expense

	sqlStatement := `SELECT id, amount FROM expenses WHERE Id = $1`
	err := er.db.QueryRow(sqlStatement, expenseID).Scan(&expense.Id, &expense.Amount)

	return expense, err
}

func (er *expenseRepository) CreateExpense(userId int, e model.Expense) (model.Expense, error) {
	var expense model.Expense
	sqlStatement := `INSERT INTO expenses(amount, expense_date, category, user_id) VALUES($1, $2, $3, $4) RETURNING amount, expense_date, category`

	err := er.db.QueryRow(sqlStatement, e.Amount, e.ExpenseDate, e.Category, userId).Scan(&expense.Amount, &expense.ExpenseDate, &expense.Category)

	return expense, err
}

func (er *expenseRepository) GetExpenses(userId int) ([]model.Expense, error) {
	var expenses []model.Expense
	sqlStatement := `SELECT id, amount, expense_date, category FROM expenses WHERE user_id = $1`

	rows, err := er.db.Query(sqlStatement, userId)
	if err != nil {
		return nil, fmt.Errorf("no expenses found: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			config.Logger.Panicw("error closing rows", "error", err)
		}
	}(rows)
	for rows.Next() {
		var expense model.Expense
		err = rows.Scan(&expense.Id, &expense.Amount, &expense.ExpenseDate, &expense.Category)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (er *expenseRepository) GetExpensesByDateRange(userID int, startDate, endDate time.Time) ([]model.Expense, error) {
	var expenses []model.Expense
	sqlStatement := `SELECT id, amount, expense_date, category FROM expenses WHERE expense_date BETWEEN $1 AND $2 AND user_id = $3`

	rows, err := er.db.Query(sqlStatement, startDate, endDate, userID)
	if err != nil {
		return nil, fmt.Errorf("no expenses found: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			config.Logger.Panicw("error closing rows", "error", err)
		}
	}(rows)
	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(&expense.Id, &expense.Amount, &expense.ExpenseDate, &expense.Category)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (er *expenseRepository) DeleteExpense(userId, expenseId int) error {
	var expense model.Expense
	sqlStatement := `DELETE FROM expenses WHERE Id = $1 AND user_id = $2`

	err := er.db.QueryRow(`SELECT id, amount, expense_date, category, user_id FROM expenses WHERE id = $1 AND user_id = $2`, expenseId, userId).Scan(&expense.Id, &expense.Amount, &expense.ExpenseDate, &expense.Category, &expense.UserId)
	if err != nil {
		return fmt.Errorf("no expense found: %v", err)
	}
	_, err = er.db.Exec(sqlStatement, expenseId, userId)

	return err
}
