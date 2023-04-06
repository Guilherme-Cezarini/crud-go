package view

import (
	"github.com/Guilherme-Cezarini/crud-go/src/controller/model/response"
	"github.com/Guilherme-Cezarini/crud-go/src/model"
)

func ConvertDomainTokenToResponse(
	tokenDomain model.TokenDomainInterface,
) response.TokenResponse {
	return response.TokenResponse{
		Token:    tokenDomain.GetToken(),
	}
}