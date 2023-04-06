package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"

)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	
	logger.Info("Init updateUser model")
	
	if(userDomain.GetPassword() != ""){
		userDomain.EncryptPassword()
	}
	
	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}
 	return nil
	
}
