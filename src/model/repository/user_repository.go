package repository

import (
	"github.com/rozzettimatheus/crud-go/src/configuration/response_error"
	"github.com/rozzettimatheus/crud-go/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) userRepository {
	return &userRepository{
		db: database,
	}
}

type UserRepository interface {
	Create(
		user model.UserDomainInterface,
	) (model.UserDomainInterface, *response_error.ResponseError)
}

type userRepository struct {
	db *mongo.Database
}
