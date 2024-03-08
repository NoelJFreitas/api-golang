package controller

import (
	"github.com/NoelJFreitas/api-golang/src/configuration/rest_errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_errors.NewBadRequestError("teste de erros da aplicação")
	c.JSON(err.Code, err)
}
