package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func InitMongo(dbName string) (MongoInstance, error) {
	mongoURI := "mongodb://localhost:27017"

	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	var db *mongo.Database = client.Database(dbName)

	if err != nil {
		return MongoInstance{}, err
	}

	instance := MongoInstance{
		Client:   client,
		Database: db,
	}

	return instance, nil
}

func (instance *MongoInstance) GetCollection(name string) *mongo.Collection {
	return instance.Database.Collection(name)
}
