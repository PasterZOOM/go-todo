package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"go-todo/internal/domain"
	"go-todo/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "aksdjhl;csxjvkjbf"
	tokenTTL   = 12 * time.Hour
	signingKey = "askjdhfalskjdcdvcas"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.UserInput) (*domain.UserResponse, error) {
	user.Password = generatePasswordHash(user.Password)

	res, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return generateUserResponse(res), nil
}

func (s *AuthService) GenerateToken(userName, password string) (string, error) {
	user, err := s.repo.GetUser(userName, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &TokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.ID,
		},
	)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(signingKey), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
