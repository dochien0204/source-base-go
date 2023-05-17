package entity

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Jti    string `json:"jti"`
}
