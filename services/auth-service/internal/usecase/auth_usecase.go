package usecase

import (
	"auth-service/internal/domain"
	jwtpkg "auth-service/internal/infrastructure/jwt"
	"auth-service/internal/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type AuthUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  repo,
		jwtSecret: jwtSecret,
	}
}

func (u *AuthUsecase) Register(email, password string) error {
	// cek apakah email sudah terdaftar
	_, err := u.userRepo.FindByEmail(email)
	if err == nil {
		// FindByEmail sukses berarti user sudah ada
		return ErrEmailAlreadyExists
	}
	if !errors.Is(err, repository.ErrUserNotFound) {
		// error lain selain "tidak ditemukan"
		return fmt.Errorf("failed to check email: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domain.User{
		Email:    email,
		Password: string(hash),
	}

	return u.userRepo.Create(user)
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			// jangan expose "email tidak ada" — gunakan pesan generik
			return "", ErrInvalidCredentials
		}
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := jwtpkg.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
