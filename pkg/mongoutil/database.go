package mongoutil

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Client
}

func New(dsn string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: client,
	}, nil
}

func ListAndCount(ctx context.Context, mc *mongo.Collection, filter bson.M, results interface{}) (int64, error) {
	total, err := mc.CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}

	cursor, err := mc.Find(ctx, filter)
	if err != nil {
		return -1, err
	}

	if err = cursor.All(ctx, results); err != nil {
		return -1, err
	}

	return total, nil
}
