package handlers

import (
	"go-web-api/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func NewPerson() *Person {
	return &Person{}
}

func (p *Person) BodyBinder(ctx *gin.Context) {
	err := ctx.BindJSON(&p)
	if err != nil {
		response := common.NewResponseModel[*Person, *Person]()
		response.ResponseCode = http.StatusBadRequest
		response.ExceptionCode = 1
		response.ExceptionModel.Message = "pasring error !"
		response.ExceptionModel.StackTrace = err.Error()
		
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, p)
}
