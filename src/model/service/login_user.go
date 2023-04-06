package service

import (
	"fmt"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/rest_err"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
	"github.com/Guilherme-Cezarini/crud-go/src/configuration/logger"

)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface,) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init loginUser model.")

	//userDomain.EncryptPassword()
	message := fmt.Sprintf("Valor do encrypt: %v", userDomain.GetPassword())

	fmt.Println(message)
	user, err := ud.userRepository.FindUserByEmail(
		userDomain.GetEmail(),
	)

	if err != nil {
		return nil, err
	}

	passwordIsValid := user.CheckPasswordHash(userDomain.GetPassword(), user.GetPassword())
	if passwordIsValid == false {
		return nil, rest_err.NewBadRequestError("Pasword or email invalid.")
	}
	
	logger.Info("loginUser service executed successfully")
	return user, nil
}