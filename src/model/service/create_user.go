package service

import (
	"fmt"

	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model")
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
