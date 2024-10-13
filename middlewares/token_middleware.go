package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenMidlleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("api-key")
		if apiKey != "" {
			ctx.Next()
		} else {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("athentication error occured !"))
		}
	}
}
