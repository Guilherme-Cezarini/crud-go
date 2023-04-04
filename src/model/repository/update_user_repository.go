package repository

import (
	"context"
	"os"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
)


func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) (*rest_err.RestErr) {
	logger.Info("Init update user repository.")
	collection_name := os.Getenv(MONGO_DB_USER)
	collection := ur.databaseConnection.Collection(collection_name)
	value := converter.ConverterDomainToEntity(userDomain)
	filter := bson.D{{Key: "_id", Value: value}}
	_, err := collection.UpdateOne(context.Background(), filter, value)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
