package controller

import (
	"net/http"
	"fmt"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/validation"
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/request"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/view"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) Insert(c *gin.Context) {
	logger.Info("Init createUser Controller.")
	var userResquet request.UserRequest

	if err := c.ShouldBindJSON(&userResquet); err != nil {
		restErr := validation.UserError(err)
		c.JSON(restErr.Code, restErr)
		logger.Error("Error trying to validate user info", err)

		return
	}

	domain := model.NewUserDomain(
		userResquet.Email,
		userResquet.Passowrd,
		userResquet.Name,
		userResquet.Age,
	)

	domainResult, err := uc.serviceInterace.CreateUser(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to create user.", err)
		return
	}

	fmt.Println(domainResult)
	logger.Info("User created success.")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))

}
