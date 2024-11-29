package service

import (
	"fmt"

	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *r.ResponseError {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.HashPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
