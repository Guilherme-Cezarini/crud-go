package service

import (
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"

)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface,) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init loginUser model.")

	userDomain.EncryptPassword()

	user, err := ud.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)

	if err != nil {
		return nil, err
	}

	logger.Info("loginUser service executed successfully")
	return user, nil
}