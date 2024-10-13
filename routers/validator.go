package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func RouteValidator(r *gin.RouterGroup) {
	r.POST("/persist", handlers.NewPersonData().Persist)
}
