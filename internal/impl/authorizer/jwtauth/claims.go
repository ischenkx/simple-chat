package jwtauth

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
}
