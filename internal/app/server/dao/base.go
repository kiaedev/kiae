package dao

import "go.mongodb.org/mongo-driver/mongo"

type Dao struct {
	collection *mongo.Collection
}

func NewDao(collection *mongo.Collection) *Dao {
	return &Dao{collection: collection}
}
