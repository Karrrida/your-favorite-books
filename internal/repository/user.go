package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(email, name, password_hash string) error {
	_, err := r.db.Exec(`
		INSERT INTO users (email, name, password_hash)
		VALUES ($1, $2, $3
	)
		`, email, name, password_hash)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {

	user := &User{}

	err := r.db.QueryRow(`
			SELECT id, email, name, created_at
			FROM users
			WHERE email = $1
		`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Created_at,
	)

	if err != nil {
		fmt.Println("Error in GetUserByEmail", err)
		return nil, err
	}

	return user, nil
}
