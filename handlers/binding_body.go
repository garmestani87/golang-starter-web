package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func NewPerson() *Person{
	return &Person{}
}

func (p *Person) BodyBinder(ctx *gin.Context) {
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError,
			errors.New("unexpected error occured !"))
	}
	ctx.JSON(http.StatusOK, p)
}
