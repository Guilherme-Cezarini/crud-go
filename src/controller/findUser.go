package controller

import (
	"net/http"
	"net/mail"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/view"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserByID controller")

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Info("Error trying to validate userId")
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.serviceInterace.FindUserByIdServices(userId)
	if err != nil {
		logger.Info("Error trying to call findUserByID services")
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByID controller")

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Info("Error trying to validate userEmail")
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.serviceInterace.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Info("Error trying to call FindUserByEmail services")
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
