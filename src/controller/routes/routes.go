package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rozzettimatheus/crud-go/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/users/:userId", userController.FindUserByID)
	r.GET("/users/email/:userId", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)
}
