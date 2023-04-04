package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository/entity/converter"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

)


func (ur *userRepository) FindUserByEmail(email string,) (model.UserDomainInterface, *rest_err.RestErr){
	logger.Info("Init find user by email repository.")
	collection_name := os.Getenv(MONGO_DB_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}

	filter := bson.D{{ Key: "email", Value: email}}
	
	err := collection.FindOne( context.Background(), filter).Decode(UserEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Info(errorMessage)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err)

		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Error("Finding user from email", err)

	return converter.ConverterEntityToDomain(*UserEntity), nil
}


func (ur *userRepository) FindUserById(id string,) (model.UserDomainInterface, *rest_err.RestErr ){
	logger.Info("Init find user by id repository.")
	collection_name := os.Getenv(MONGO_DB_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}
	userId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{ Key: "_id", Value: userId}}	
	err := collection.FindOne(context.Background(), filter).Decode(UserEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id)
			logger.Info(errorMessage)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err)

		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Error("Finding user from id", err)

	return converter.ConverterEntityToDomain(*UserEntity), nil
}