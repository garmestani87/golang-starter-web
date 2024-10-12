package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromParam(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	name, _ := ctx.Params.Get("name")

	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
