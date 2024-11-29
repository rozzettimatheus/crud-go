package service

import (
	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUser(id string) (*model.UserDomainInterface, *r.ResponseError) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))
	return nil, nil
}
