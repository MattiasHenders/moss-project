package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthToken struct {
	AccessToken string `json:"accessToken"`
	Expires     int64  `json:"expires"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}
