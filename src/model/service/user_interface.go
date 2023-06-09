package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(id string,) (model.UserDomainInterface, *rest_err.RestErr )
	FindUserByEmailServices(email string,) (model.UserDomainInterface, *rest_err.RestErr )
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr )
	DeleteUser(string) *rest_err.RestErr
}
