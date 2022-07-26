package repository

import (
	"database/sql"
	"fmt"
	"reimbursement_backend/config"
	"reimbursement_backend/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
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
	var role string
	sqlStatement := `INSERT INTO users(name, email, role) VALUES($1, $2, $3) RETURNING id, name, email, role`
	adminEmails := config.Config.Email.AdminEmails
	caEmail := config.Config.Email.CAEmail

	for _, email := range adminEmails {
		if user.Email == email {
			role = "admin"
			break
		}
	}
	if role == "" {
		if user.Email == caEmail {
			role = "ca"
		} else {
			role = "employee"
		}
	}
	err := ur.db.QueryRow(sqlStatement, user.Name, user.Email, role).Scan(&createdUser.Id, &createdUser.Name, &createdUser.Email, &createdUser.Role)

	return createdUser, err
}

func (ur *userRepository) GetUserByEmail(email string) (model.User, error) {
	var foundUser model.User
	sqlStatement := `SELECT id, name, email, role FROM users where email = $1`

	err := ur.db.QueryRow(sqlStatement, email).Scan(&foundUser.Id, &foundUser.Name, &foundUser.Email, &foundUser.Role)

	return foundUser, err
}

func (ur *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	sqlStatement := `SELECT id, name, email, role FROM users`

	rows, err := ur.db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("repo: could not fetch users: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			config.Logger.Panicw("error closing rows", "error", err)
		}
	}(rows)
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("repo: error scanning users: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
