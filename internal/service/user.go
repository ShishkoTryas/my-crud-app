package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"my-crud-app/internal/domain"
	"strconv"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, email, password string) (domain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type Users struct {
	repo   UserRepository
	hasher PasswordHasher

	hmacSecret []byte
}

func New(repo UserRepository, hasher PasswordHasher, secret []byte) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
	}
}

func (u *Users) SignUp(ctx context.Context, inp domain.SignUpUser) error {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}
	return u.repo.Create(ctx, user)
}

func (u *Users) SignIn(ctx context.Context, inp domain.SignInUser) (string, error) {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return "", err
	}

	user, err := u.repo.GetUser(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("User not found")
		}

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})

	return token.SignedString(u.hmacSecret)

}

func (u *Users) ParseToken(ctx context.Context, token string) (int64, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return u.hmacSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return int64(id), nil
}
