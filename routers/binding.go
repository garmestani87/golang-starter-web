package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-api/handlers"
)

func BindingRouter(r *gin.RouterGroup) {
	//header
	r.POST("/header_1", handlers.NewUserInfo().ReadFromHeader)
	r.POST("/header_2", handlers.NewUserInfo().ReadAndBindHeader)

	//query
	r.POST("/query_1", handlers.ReadFromQuery)
	r.POST("/query_2", handlers.ReadArrayFromQuery)

	//path variable
	r.GET("/path_variable_1/:id/:name", handlers.ReadFromParam)

	//body
	r.POST("/body", handlers.NewPerson().BodyBinder)
	
	//form
	r.POST("/form", handlers.NewFormData().FormBinder)
}
