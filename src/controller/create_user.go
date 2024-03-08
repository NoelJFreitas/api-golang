package controller

import (
	"fmt"

	"github.com/NoelJFreitas/api-golang/src/configuration/validation"
	"github.com/NoelJFreitas/api-golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//converte o json e joga na variável userRequest
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUseError(err)
		c.JSON(restErr.Code, restErr)
		fmt.Println(err.Error())
		return
	}

	fmt.Println(userRequest)
}
