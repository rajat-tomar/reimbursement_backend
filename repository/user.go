package repository

import (
	"database/sql"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
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
	sqlStatement := `INSERT INTO users(name, email, role) VALUES($1, $2, $3) RETURNING id, name, email, role`

	err := ur.db.QueryRow(sqlStatement, user.Name, user.Email, "employee").Scan(&user.Id, &user.Name, &user.Email, &user.Role)

	return user, err
}

func (ur *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	sqlStatement := `SELECT * FROM users where email = $1`

	err := ur.db.QueryRow(sqlStatement, email).Scan(&user.Name, &user.Email, &user.Role)

	return user, err
}
