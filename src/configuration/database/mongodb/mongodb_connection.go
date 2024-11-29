package mongodb

import (
	"context"
	"os"

	"github.com/rozzettimatheus/crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	DATABASE_URL = "DATABASE_URL"
	USERS_DB     = "USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(DATABASE_URL)
	client, conn_err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))
	if conn_err != nil {
		return nil, conn_err
	}
	if ping_err := client.Ping(ctx, nil); ping_err != nil {
		return nil, ping_err
	}
	logger.Info("Database successfully connected")
	return client.Database(USERS_DB), nil
}
