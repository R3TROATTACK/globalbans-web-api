package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"insanitygaming.net/bans/services/logger"
)

var client *mongo.Client

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}

	db := client.Database("globalbans")
	err = db.CreateCollection(ctx, "admins")
	if err != nil {
		logger.Logger().Warn(err)
	}
	err = db.CreateCollection(ctx, "bans")
	if err != nil {
		logger.Logger().Warn(err)
	}
	err = db.CreateCollection(ctx, "servers")
	if err != nil {
		logger.Logger().Warn(err)
	}
	err = db.CreateCollection(ctx, "servergroups")
	if err != nil {
		logger.Logger().Warn(err)
	}
	err = db.CreateCollection(ctx, "webgroups")
	if err != nil {
		logger.Logger().Warn(err)
	}
	err = db.CreateCollection(ctx, "admingroups")
	if err != nil {
		logger.Logger().Warn(err)
	}

}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database("globalbans")
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	logger.Logger().Info("Disconnected from MongoDB")
}
