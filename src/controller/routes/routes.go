package routes

import (
	"github.com/NoelJFreitas/api-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/id/:userId", controller.FindUserByID)
	r.GET("/user/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
