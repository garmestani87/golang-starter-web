package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonData struct {
	Name         string `json:"name" binding:"required,max=10"`
	Age          int    `json:"age" binding:"required,gte=13"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile"`
}

func NewPersonData() *PersonData {
	return &PersonData{}
}

func (p *PersonData) Persist(ctx *gin.Context) {
	if err := ctx.ShouldBindBodyWithJSON(&p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"description": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"name":   p.Name,
		"mobile": p.MobileNumber,
	})
}
