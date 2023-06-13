package mongo

import (
	"context"
	"errors"
	"logging-service/internal/database"
	"logging-service/internal/repository"
	"logging-service/pkg/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepositoryImpl struct {
	instance database.MongoInstance
}

func New(instance database.MongoInstance) repository.LogRepository {
	return &LogRepositoryImpl{
		instance: instance,
	}
}

func (repo *LogRepositoryImpl) Write(logInput model.LogInput) (*model.Log, error) {
	var blog model.Log = model.Log{
		UserID:    logInput.UserID,
		BookID:    logInput.BookID,
		BookTitle: logInput.BookTitle,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = repo.instance.GetCollection("logs")

	result, err := collection.InsertOne(context.TODO(), blog)

	if err != nil {
		return &model.Log{}, errors.New("create log failed")
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(context.TODO(), filter)

	var createdLog *model.Log = &model.Log{}

	createdRecord.Decode(createdLog)

	return createdLog, nil
}
