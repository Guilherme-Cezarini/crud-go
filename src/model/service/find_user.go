package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
)

func (ud *userDomainService) FindUserByIdServices(id string, ) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByIdServices.")
	return ud.userRepository.FindUserById(id)
}


func (ud *userDomainService) FindUserByEmailServices(email string, ) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailServices model")

	return ud.userRepository.FindUserByEmail(email)
}
