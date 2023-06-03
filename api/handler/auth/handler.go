package handler

import (
	"source-base-go/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, authService user.UseCase) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/login", func(ctx *gin.Context) {
			login(ctx, authService)
		})
	}
}
