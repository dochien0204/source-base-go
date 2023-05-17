package util

import (
	"errors"
	"source-base-go/config"
	"source-base-go/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func keyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New(config.UNAUTHORIZED)
		}

		return []byte(key), nil
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	//Get Authorization Header
	authHeader := ctx.GetHeader(config.GetString("jwt.header"))
	if len(authHeader) <= len(config.GetString("jwt.schema")) {
		return "", errors.New(config.UNAUTHORIZED)
	}

	//Get Token from Authorization Header ("jwt.schema" = Bearer)
	tokenString := authHeader[len(config.GetString("jwt.schema"))+1:]

	return tokenString, nil
}

// Validate Token
func ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, keyFunc(config.GetString("jwt.secretKey")))
}

func ParseAccessToken(token string) (*entity.TokenClaims, error) {
	claims := entity.TokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc(config.GetString("jwt.secretKey")))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
