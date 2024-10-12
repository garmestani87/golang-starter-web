package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id   int
	Name string
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (u *UserInfo) ReadFromHeader(ctx *gin.Context) {
	apiKey := ctx.GetHeader("api-key")
	ctx.JSON(http.StatusOK, gin.H{
		"apiKey": apiKey,
		"id":     10,
		"name":   "jafar",
	})
}

func (u *UserInfo) ReadFromBindHeader(ctx *gin.Context) {
	ctx.BindHeader(&u)
	ctx.JSON(http.StatusOK, gin.H{
		"id":   u.Id,
		"name": u.Name,
	})
}
