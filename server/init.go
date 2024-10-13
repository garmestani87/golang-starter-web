package server

import (
	"go-web-api/config"
	"go-web-api/middlewares"
	"go-web-api/routers"
	"go-web-api/validators"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TokenMidlleware(), middlewares.AclMiddleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", validators.MobileValidator)
	}
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			routers.HealthRouter(v1.Group("/health"))
			routers.UsersRouter(v1.Group("/users"))
			routers.BindingRouter(v1.Group("/binding"))
			routers.RouteValidator(v1.Group("/validate"))
		}

	}

	os.Setenv("APP_ENV", "dev")
	cfg := config.GetConfig()

	err := r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
