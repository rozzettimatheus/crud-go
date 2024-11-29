package service

import (
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *r.ResponseError
	UpdateUser(string, model.UserDomainInterface) *r.ResponseError
	FindUser(string) (*model.UserDomainInterface, *r.ResponseError)
	DeleteUser(string) *r.ResponseError
}
