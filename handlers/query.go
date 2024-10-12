package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromQuery(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func ReadArrayFromQuery(ctx *gin.Context) {
	names := ctx.QueryArray("name")
	id := ctx.Query("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"names": names,
	})
}
