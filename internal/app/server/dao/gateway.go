package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/gateway"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Gateway struct {
	*Dao
}

func NewGateway(db *mongo.Database) *Gateway {
	return &Gateway{
		Dao: NewDao(db.Collection("clusters")),
	}
}

func (p *Gateway) Get(ctx context.Context, id string) (*gateway.Gateway, error) {
	var proj gateway.Gateway
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *Gateway) List(ctx context.Context, m bson.M) ([]*gateway.Gateway, int64, error) {
	var results []*gateway.Gateway
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *Gateway) Create(ctx context.Context, in *gateway.Gateway) (*gateway.Gateway, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *Gateway) Update(ctx context.Context, in *gateway.Gateway) (*gateway.Gateway, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
