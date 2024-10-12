package routers

import (
	"go-web-api/handlers"

	"github.com/gin-gonic/gin"
)

func BindHeaderRouter(r *gin.RouterGroup) {
	r.POST("/header_1", handlers.NewUserInfo().ReadFromHeader)
	r.POST("/header_2", handlers.NewUserInfo().ReadAndBindHeader)

	r.POST("/query_1", handlers.ReadFromQuery)
	r.POST("/query_2", handlers.ReadArrayFromQuery)
	
	r.GET("/param_1", handlers.ReadFromParam)
}
