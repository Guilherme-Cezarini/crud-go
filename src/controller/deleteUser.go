package controller

import (
	"net/http"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) Delete(c *gin.Context) {

	logger.Info("Init Delete Controller.")
	
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.serviceInterace.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to delete user.", err)
		return
	}

	logger.Info("delete success.")
	c.Status(http.StatusOK)

}
