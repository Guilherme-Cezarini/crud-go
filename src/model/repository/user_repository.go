package repository

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_DB_USER = "MONGODB_USER_DB"
)


func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(id string,) (model.UserDomainInterface, *rest_err.RestErr )
	FindUserByEmail(email string,) (model.UserDomainInterface, *rest_err.RestErr )
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string,) *rest_err.RestErr
}
