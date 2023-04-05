package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model")
	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil
}
