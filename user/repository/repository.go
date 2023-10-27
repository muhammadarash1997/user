package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhammadarash1997/user/user/model"

	"fmt"
)

var (
	insertUser = `
		INSERT INTO users(username, email, password_hash) VALUES (?,?,?)
	`
	selectUserById = `
		SELECT * FROM users WHERE id = ?
	`
	selectUserByEmail = `
		SELECT * FROM users WHERE email = ?
	`
)

type Repository interface {
	CreateUser(model.User) error
	GetUserById(uint) (model.User, error)
	GetUserByEmail(string) (model.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user model.User) error {
	valueArgs := []interface{}{}
	valueArgs = append(valueArgs, user.Username)
	valueArgs = append(valueArgs, user.Email)
	valueArgs = append(valueArgs, user.Password)

	_, err := r.db.Exec(insertUser, valueArgs...)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}

func (r *repository) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(selectUserByEmail, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("r.db.QueryRow: %v", err)
	}

	return user, nil
}

func (r *repository) GetUserById(id uint) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(selectUserById, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("r.db.QueryRow: %v", err)
	}

	return user, nil
}
