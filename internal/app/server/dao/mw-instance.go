package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MiddlewareInstance struct {
	*Dao
}

func NewMiddlewareInstanceDao(db *mongo.Database) *MiddlewareInstance {
	return &MiddlewareInstance{
		Dao: NewDao(db.Collection("middleware-instances")),
	}
}

func (p *MiddlewareInstance) Get(ctx context.Context, id string) (*middleware.Instance, error) {
	var res middleware.Instance
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *MiddlewareInstance) GetDefault(ctx context.Context, mType string, env string) (*middleware.Instance, error) {
	var res middleware.Instance
	if err := p.collection.FindOne(ctx, bson.M{"type": mType, "env": env, "status": kiae.OpStatus_OP_STATUS_ENABLED}).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *MiddlewareInstance) FindByAppid(ctx context.Context, mType string, appid string) (*middleware.Instance, error) {
	var res middleware.Instance
	if err := p.collection.FindOne(ctx, bson.M{"type": mType, "bindings": bson.M{"$in": appid}}).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *MiddlewareInstance) List(ctx context.Context, m bson.M) ([]*middleware.Instance, int64, error) {
	var results []*middleware.Instance
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *MiddlewareInstance) Create(ctx context.Context, in *middleware.Instance) (*middleware.Instance, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	if err == nil {
		in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	}
	return in, err
}

func (p *MiddlewareInstance) Update(ctx context.Context, in *middleware.Instance) (*middleware.Instance, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	in.Id = oid.Hex()
	return in, err
}
