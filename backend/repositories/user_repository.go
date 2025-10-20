package repositories

import (
	"banque-app/backend/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)", user.Name, user.Email, user.PasswordHash)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, password_hash FROM users WHERE email = ?", email)
	var u models.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
		return nil, err
	}
	return &u, nil
}
