package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MiddlewareClaim struct {
	*Dao
}

func NewMiddlewareClaimDao(db *mongo.Database) *MiddlewareClaim {
	return &MiddlewareClaim{
		Dao: NewDao(db.Collection("middleware-claims")),
	}
}

func (p *MiddlewareClaim) Get(ctx context.Context, id string) (*middleware.Claim, error) {
	var res middleware.Claim
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *MiddlewareClaim) List(ctx context.Context, m bson.M) ([]*middleware.Claim, int64, error) {
	var results []*middleware.Claim
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *MiddlewareClaim) Create(ctx context.Context, in *middleware.Claim) (*middleware.Claim, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	if err == nil {
		in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	}
	return in, err
}

func (p *MiddlewareClaim) Update(ctx context.Context, in *middleware.Claim) (*middleware.Claim, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	in.Id = oid.Hex()
	return in, err
}
