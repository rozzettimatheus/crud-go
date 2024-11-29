package view

import (
	"github.com/rozzettimatheus/crud-go/src/controller/model/response"
	"github.com/rozzettimatheus/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetEmail(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
