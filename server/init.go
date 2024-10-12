package server

import (
	"go-web-api/config"
	"go-web-api/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	health := r.Group("/health")
	routers.Route(health)

	os.Setenv("APP_ENV", "dev")
	cfg := config.GetConfig()

	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
