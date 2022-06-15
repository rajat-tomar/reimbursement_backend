package repository

import (
	"gorm.io/gorm"
	"reimbursement_backend/model"
)

type ExpenseRepository interface {
	Create(expense model.Expense) (int, error)
	GetById(expenseID int) (float64, error)
	GetAll() ([]model.Expense, error)
}

type expenseRepository struct {
	db *gorm.DB
}

func (er *expenseRepository) Create(expense model.Expense) (int, error) {
	tx := er.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx = tx.Create(&expense)
	if tx.Error != nil {
		tx.Rollback()
		return 0, tx.Error
	}

	err := tx.Commit().Error
	if err != nil {
		return 0, err
	}
	return expense.ID, nil
}
