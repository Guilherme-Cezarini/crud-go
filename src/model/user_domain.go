package model

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string

	SetID(string)

	EncryptPassword()
	CheckPasswordHash(string, string) bool
}

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(ud.password), 14)
	ud.password = string(bytes)
}

func (ud *userDomain) CheckPasswordHash(password, hash string) bool {
	message := fmt.Sprintf("password: %v - hash: %v", password, hash)
	fmt.Println(message)
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

