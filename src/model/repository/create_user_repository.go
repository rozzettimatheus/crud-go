package repository

import (
	"github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
)

func (r *userRepository) Create(
	user model.UserDomainInterface,
) (model.UserDomainInterface, *response_error.ResponseError) {

}
