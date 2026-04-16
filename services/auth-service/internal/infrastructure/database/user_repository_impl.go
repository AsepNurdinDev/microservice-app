package database

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"database/sql"
	"errors"
	"fmt"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *domain.User) error {
	query := `INSERT INTO users(email, password, created_at) VALUES($1, $2, NOW())`
	_, err := r.db.Exec(query, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, email, password, created_at FROM users WHERE email=$1`

	user := &domain.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}