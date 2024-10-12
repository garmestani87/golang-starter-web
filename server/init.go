package server

import (
	"go-web-api/routers"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	health := r.Group("/health")
	routers.Route(health)

	err := r.Run(":30080")
	if err != nil {
		panic(err)
	}
}
