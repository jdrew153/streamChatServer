package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once
var clientInstanceError error

type Collection string

const (
	UsersCollection Collection = "users"
)

const (
	uri      = "mongodb+srv://jdrew153:Katlyn29@cluster0.mfo1oqg.mongodb.net/?retryWrites=true&w=majority"
	Database = "streamMessaging"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
