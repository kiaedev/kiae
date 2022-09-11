package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/route"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RouteDao struct {
	*Dao
}

func NewRouteDao(db *mongo.Database) *RouteDao {
	return &RouteDao{
		Dao: NewDao(db.Collection("routes")),
	}
}

func (p *RouteDao) Get(ctx context.Context, id string) (*route.Route, error) {
	var proj route.Route
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *RouteDao) List(ctx context.Context, m bson.M) ([]*route.Route, int64, error) {
	var results []*route.Route
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *RouteDao) Create(ctx context.Context, in *route.Route) (*route.Route, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	if err == nil {
		in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	}
	return in, err
}

func (p *RouteDao) Update(ctx context.Context, in *route.Route) (*route.Route, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	in.Id = oid.Hex()
	return in, err
}
