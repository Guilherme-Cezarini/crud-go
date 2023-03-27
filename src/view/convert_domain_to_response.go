package view

import (
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/response"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id:    0,
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
