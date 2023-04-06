package service

import (
	"fmt"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"

)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model")

	oldUSer, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if(oldUSer != nil) {
		return nil, rest_err.NewInternalServerError("Email in use.")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	
	return userDomainRepository, nil
}
