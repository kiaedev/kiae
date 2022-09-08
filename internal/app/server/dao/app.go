package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppDao struct {
	*Dao
}

func NewApp(db *mongo.Database) *AppDao {
	return &AppDao{
		Dao: NewDao(db.Collection("apps")),
	}
}

func (p *AppDao) Get(ctx context.Context, id string) (*app.Application, error) {
	var proj app.Application
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *AppDao) List(ctx context.Context, m bson.M) ([]*app.Application, int64, error) {
	var results []*app.Application
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *AppDao) Create(ctx context.Context, in *app.Application) (*app.Application, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *AppDao) Update(ctx context.Context, in *app.Application) (*app.Application, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
