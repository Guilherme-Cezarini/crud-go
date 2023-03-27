package main

import (
	"context"
	"log"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/database/mongodb"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/controller"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository"
	"github.com/Guilherme-Cezarini/crud-go/src/model/service"

	"github.com/Guilherme-Cezarini/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Erro trying to connect to database, error=%s", err.Error())
	}
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
