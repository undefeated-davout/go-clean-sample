package repository

import (
	"database/sql"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MySQLUserRepository struct {
	conn *sql.DB
}

func NewMySQLUserRepository(conn *sql.DB) domain.UserRepository {
	return &MySQLUserRepository{conn: conn}
}

func (r *MySQLUserRepository) Save(user domain.User) error {
	_, err := r.conn.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}

func (r *MySQLUserRepository) GetByID(id int) (*domain.User, error) {
	row := r.conn.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*domain.User, error) {
	row := r.conn.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) Delete(id int) error {
	_, err := r.conn.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
