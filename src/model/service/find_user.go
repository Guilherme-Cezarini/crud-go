package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"

)

func (ud *userDomainService) FindUserByIdServices(id string,) (model.UserDomainInterface, *rest_err.RestErr ) {
	logger.Info("Init find user by id service.")
	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string,) (model.UserDomainInterface, *rest_err.RestErr ){
	logger.Info("Init find user email id service.")
	return ud.userRepository.FindUserByEmail(email)
}
