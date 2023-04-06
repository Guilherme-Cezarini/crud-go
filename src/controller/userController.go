package controller

import (
	"github.com/Guilherme-Cezarini/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterace service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		serviceInterace: serviceInterace,
	}
}

type UserControllerInterface interface {
	Update(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Login(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	serviceInterace service.UserDomainService
}
