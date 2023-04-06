package controller

import (
	"context"
	"bytes"
	"os"
	"fmt"
	//"strings"
	//"io/ioutil"

	//"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/database/mongodb"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository"
	"github.com/Guilherme-Cezarini/crud-go/src/model/service"
	"testing"
	//"github.com/Guilherme-Cezarini/crud-go/src/view"
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	//"github.com/Guilherme-Cezarini/crud-go/src/controller/middleware"

)

type Reader interface {
	Read(buf []byte) (n int, err error)
}

func SetUpRouter() *gin.Engine{
    router := gin.Default()

	

    return router
}

func router() *gin.Engine {
	router := gin.Default()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		fmt.Printf("Erro trying to connect to database, error=%s", err.Error())
	}
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := NewUserControllerInterface(service)
	//router.Use(middleware.AuthUser())
	router.POST("/create", userController.Insert)

	return router
}

func TestInsert(t *testing.T) {
	
	body := []byte(`{
		"email": "teste@teste.com",
		"name":   "teste",
		"password": "123456@",
		"age":  25
	}`)
	
	

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	tokenInConfig := os.Getenv("BEARER_TOKEN")
	bearerToken := fmt.Sprintf("Bearer %v", tokenInConfig)
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
    router().ServeHTTP(w, req)
	message := fmt.Sprintf("%v",w)
	t.Errorf(message)
	/*user := fmt.Sprintf("%v",domain);
	domainView := fmt.Sprintf("%v",jsonValue);*/
	//fmt.Println(w, domainView, user)
	//defer response.Body.Close()
	assert.Equal(t, http.StatusOK, w.Code)
}