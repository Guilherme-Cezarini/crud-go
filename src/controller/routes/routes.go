package routes

import (
	"github.com/Guilherme-Cezarini/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/create", userController.Insert)
	r.PUT("/update/:userId", userController.Update)
	r.DELETE("/delete/:userId", userController.Delete)
}
