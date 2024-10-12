package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.GET("/ping", handlers.New().Handle)
}
