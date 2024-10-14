package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AclMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader("api-key")

		if apikey != "01dd0893-f81f-43af-8ebc-94948fd7aae0" {
			ctx.AbortWithError(http.StatusForbidden, errors.New("you can not consume this service"))
		}
		ctx.Next()
	}
}
