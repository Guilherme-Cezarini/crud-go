package model


type TokenDomainInterface interface {
	GetToken() string
}

type tokenDomain struct {
	token	string
}


func NewTokenDomain(
	token string,
	
) TokenDomainInterface {
	return &tokenDomain{
		token:    token,
	}
}

func (td *tokenDomain)GetToken() string {
	return td.token
}
