package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct{}

func NewUsers() *Users {
	return &Users{}
}

func (u *Users) GetUserById(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": "jafar",
			"age":  28,
		})
	}

}
