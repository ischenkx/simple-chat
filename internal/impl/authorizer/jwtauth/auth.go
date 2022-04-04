package jwtauth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/ischenkx/vk-test-task/internal/app/security"
	"log"
	"time"
)

type Auth struct {
	expirationTime time.Duration
	key            []byte
}

func (a *Auth) Verify(ctx context.Context, tokenString string) (userId string, err error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.key, nil
	})

	if err != nil {
		log.Println(err)
		return "", errors.New("failed to parse token: " + err.Error())
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	if time.Now().After(time.Unix(claims.ExpiresAt, 0)) {
		return "", errors.New("expired token")
	}

	return claims.Id, nil
}

func (a *Auth) GenerateToken(ctx context.Context, userId string) (token string, err error) {
	var claims Claims
	claims.Id = userId
	claims.ExpiresAt = time.Now().Add(a.expirationTime).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(a.key)
}

func New(key []byte, expirationTime time.Duration) security.Authorizer {
	return &Auth{
		expirationTime: expirationTime,
		key:            key,
	}
}
