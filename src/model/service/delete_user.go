package service

import (
	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUser(id string) *r.ResponseError {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
