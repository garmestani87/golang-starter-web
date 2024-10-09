package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong !!!")
	})

	err := r.Run(":30080")
	if err != nil {
		panic(err)
	}
}
