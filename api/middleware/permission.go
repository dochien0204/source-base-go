package middleware

import (
	"net/http"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permisstion define.Permisstion) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Get user data from gin Context
		value, exists := ctx.Get("userData")
		if !exists {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}

		//cast user data
		userData, ok := value.(*util.UserData)
		if !ok {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}

		//Check permission user
		if !util.InArray(string(permisstion), userData.ListRole) {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}

		ctx.Next()
	}
}
