package repository

import "github.com/thitiphongD/go-backend/internal/auth/domain"

type UserRepository interface {
	GetByEmail(email string) (*domain.User, error)
}
