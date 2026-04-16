package repository

import (
	"auth-service/internal/domain"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}