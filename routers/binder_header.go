package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func BindHeaderRouter(r *gin.RouterGroup) {
	r.POST("/read-key", handlers.NewUserInfo().ReadFromHeader)
	r.POST("/read-bind", handlers.NewUserInfo().ReadFromBindHeader)
}
