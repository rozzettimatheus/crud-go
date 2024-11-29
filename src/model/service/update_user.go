package service

import (
	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *r.ResponseError {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))
	return nil
}
