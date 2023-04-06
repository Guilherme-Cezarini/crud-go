package routes

import (
	"github.com/Guilherme-Cezarini/crud-go/src/controller"
	"github.com/Guilherme-Cezarini/crud-go/src/controller/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	
	r.POST("/login", userController.Login)


	r.Use(middleware.AuthUser())
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/create", userController.Insert)
	r.PUT("/update/:userId", userController.Update)
	r.DELETE("/delete/:userId", userController.Delete)
}
