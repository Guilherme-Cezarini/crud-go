package routes

import (
	"github.com/Guilherme-Cezarini/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/create", controller.Insert)
	r.PUT("/update/:userId", controller.Update)
	r.DELETE("/delete/:userId", controller.Delete)
}
