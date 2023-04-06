package middleware

import(
    "strings"
    "os"
    "github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
    "github.com/go-playground/validator/v10"
    "github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"


    "github.com/gin-gonic/gin"
)


type authHeader struct {
    IDToken string `header:"Authorization"`
}


func AuthUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        logger.Info("Init auth middleware")    

        h := authHeader{}
        var tokenInConfig string

        if err := c.ShouldBindHeader(&h); err != nil {
            if _, ok := err.(validator.ValidationErrors); ok {
                logger.Info("Error on validation header auth.")    

                rest_err := rest_err.NewInternalServerError(err.Error())
                c.JSON(rest_err.Code, rest_err)
                c.Abort()
                return
            }

            rest_err := rest_err.NewInternalServerError(err.Error())
            logger.Info("Error on validation header auth.")    

            c.JSON(rest_err.Code, rest_err)
            c.Abort()
            return
        }
        tokenInConfig = os.Getenv("BEARER_TOKEN")
       
        if tokenInConfig == "" {
            logger.Info("Token is empty in .env file.")    
            rest_err := rest_err.NewInternalServerError("No token for validate.")
            c.JSON(rest_err.Code, rest_err)
            c.Abort()
            return
        }

        idTokenHeader := strings.Split(h.IDToken, "Bearer ")
        if (len(idTokenHeader) < 2 || idTokenHeader[1] != tokenInConfig) {
            logger.Info("Invalid token")    
            rest_err := rest_err.NewForbiddenError("Invalid Token")
            c.JSON(rest_err.Code, rest_err)
            c.Abort()
            return
        }

        logger.Info("Token is valid.")    
        c.Next()
    }
}