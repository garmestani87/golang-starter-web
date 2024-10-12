package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.RouterGroup) {
	r.GET("/find-by-id/:id", handlers.NewUsers().GetUserById)
}
