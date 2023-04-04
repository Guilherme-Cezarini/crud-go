package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"net/mail"
	"net/http"
	"github.com/Guilherme-Cezarini/crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

<<<<<<< Updated upstream
func FindUserById(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {
=======
func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FondeUserById controller.")

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Info("Erro trying validate userId param.")

		errorMessage := rest_err.NewBadRequestError(
			"userID is not a valid id.",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.serviceInterace.FindUserByIdServices(userId)
	if err != nil {
		logger.Info("Erro trying find userId data.")
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

	logger.Info("FindUserById success.")

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FondeUserByEmail controller.")

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Info("Erro trying validate userEmail param.")

		errorMessage := rest_err.NewBadRequestError(
			"userEmail is not a valid email.",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.serviceInterace.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Info("Erro trying find userEmail data.")
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
>>>>>>> Stashed changes

	logger.Info("FindUserByEmail success.")
}
