package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("supersecretkey")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
