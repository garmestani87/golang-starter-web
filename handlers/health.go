package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func New() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong !!!")
}
