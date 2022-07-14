package repository

import (
	"database/sql"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: config.GetDb(),
	}
}

func (ur *userRepository) CreateUser(user model.User) (model.User, error) {
	var createdUser model.User
	sqlStatement := `INSERT INTO users(name, email, role) VALUES($1, $2, $3) RETURNING id, name, email, role`

	err := ur.db.QueryRow(sqlStatement, user.Name, user.Email, "employee").Scan(&createdUser.Id, &createdUser.Name, &createdUser.Email, &createdUser.Role)

	return createdUser, err
}

func (ur *userRepository) FindByEmail(email string) (model.User, error) {
	var foundUser model.User
	sqlStatement := `SELECT id, name, email, role FROM users where email = $1`

	err := ur.db.QueryRow(sqlStatement, email).Scan(&foundUser.Id, &foundUser.Name, &foundUser.Email, &foundUser.Role)

	return foundUser, err
}
