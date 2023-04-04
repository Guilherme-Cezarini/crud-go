package controller

import (
	"net/http"
	"fmt"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/validation"
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/request"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) Update(c *gin.Context) {

	logger.Info("Init createUser Controller.")
	var userResquet request.UserUpdateRequest
	
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	if err := c.ShouldBindJSON(&userResquet); err != nil {
		restErr := validation.UserError(err)
		c.JSON(restErr.Code, restErr)
		logger.Error("Error trying to validate user info", err)

		return
	}

	domain := model.NewUserDomain(
		userResquet.Email,
		userResquet.Name,
		userResquet.Passowrd,
		userResquet.Age,
	)

	err := uc.serviceInterace.UpdateUser(userId,domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to update user.", err)
		return
	}

	fmt.Println(domain)
	logger.Info("update success.")
	c.Status(http.StatusOK)

}
