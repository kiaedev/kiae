package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	collection *mongo.Collection
}

func NewDao(collection *mongo.Collection) *Dao {
	return &Dao{collection: collection}
}

func (p *Dao) Delete(ctx context.Context, id string) (err error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err = p.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return
}
