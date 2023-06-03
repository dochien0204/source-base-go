package middleware

import (
	"net/http"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"

	"github.com/gin-gonic/gin"
)

func JWTVerifyMiddleware(verifier util.JWTVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetToken(ctx)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		isVerified, userData, err := verifier.VerifyToken(token)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		if userData.Status != define.ACTIVE {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}

		if isVerified {
			ctx.Set("userData", userData)
			ctx.Next()
			return
		} else {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}
	}
}
