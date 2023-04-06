package repository

import (
	"context"
	"os"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (ur *userRepository) DeleteUser(userId string) (*rest_err.RestErr) {
	logger.Info("Init delete user repository.")
	collection_name := os.Getenv(MONGO_DB_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
