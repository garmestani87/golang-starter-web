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

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			health := v1.Group("/health")
			routers.HealthRouter(health)

			users := v1.Group("/users")
			routers.UsersRouter(users)

			header := v1.Group("/header")
			routers.BindHeaderRouter(header)
		}

	}

	os.Setenv("APP_ENV", "dev")
	cfg := config.GetConfig()

	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
