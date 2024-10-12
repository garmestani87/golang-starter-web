package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	r.GET("/ping", handlers.NewHealth().Check)
}
