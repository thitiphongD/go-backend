// usecase/auth_usecase.go
package usecase

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/thitiphongD/go-backend/internal/auth/domain"
	"github.com/thitiphongD/go-backend/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("daew")

type AuthUsecase interface {
	SignIn(email, password string) (*domain.User, string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(ur repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: ur,
	}
}

func (a *authUsecase) SignIn(email, password string) (*domain.User, string, error) {
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
