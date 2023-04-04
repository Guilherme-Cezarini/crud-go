package controller

import (
	"net/http"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/validation"
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/request"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func Insert(c *gin.Context) {
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

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to create user.", err)
		return
	}

	logger.Info("User created success.")
	c.String(http.StatusOK, "")

}
