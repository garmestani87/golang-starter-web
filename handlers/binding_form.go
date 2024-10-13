package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Id   int
	Name string
	Age  int
}

func NewFormData() *FormData {
	return &FormData{}
}

func (f *FormData) FormBinder(ctx *gin.Context) {
	err := ctx.ShouldBind(&f)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest,
			errors.New("bad request !"))
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "reza",
		"age":  10,
	})
}
