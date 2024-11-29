package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	"github.com/rozzettimatheus/crud-go/src/configuration/validation"
	"github.com/rozzettimatheus/crud-go/src/controller/model/request"
	"github.com/rozzettimatheus/crud-go/src/model"
	"github.com/rozzettimatheus/crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userReq request.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		responseError := validation.ValidateUserErr(err)
		c.JSON(responseError.Code, responseError)
		return
	}
	domain := model.NewUserDomain(
		userReq.Email,
		userReq.Password,
		userReq.Name,
		userReq.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}
	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
