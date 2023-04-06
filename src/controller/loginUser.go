package controller


import(
	"net/http"
	"os"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/validation"
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/request"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/view"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) Login(c *gin.Context) {
	logger.Info("Init loginUser controller")
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Info("Error trying to validate user info")
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)


	_, err := uc.serviceInterace.LoginUserServices(domain)
	if err != nil {
		logger.Info("Error trying to call loginUser service")
		c.JSON(err.Code, err)
		return
	}

	tokenDomain := model.NewTokenDomain(
		os.Getenv("BEARER_TOKEN"),
	)

	logger.Info("loginUser controller executed successfully")

	c.JSON(http.StatusOK, view.ConvertDomainTokenToResponse(
		tokenDomain,
	))
}