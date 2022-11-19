package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ImageRegistryDao struct {
	*Dao
}

func NewRegistryDao(db *mongo.Database) *ImageRegistryDao {
	return &ImageRegistryDao{
		Dao: NewDao(db.Collection("registry")),
	}
}

func (p *ImageRegistryDao) Get(ctx context.Context, id string) (*image.Registry, error) {
	var rt image.Registry
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ImageRegistryDao) GetByName(ctx context.Context, name string) (*image.Registry, error) {
	var rt image.Registry
	if err := p.collection.FindOne(ctx, bson.M{"name": name}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ImageRegistryDao) List(ctx context.Context, m bson.M) ([]*image.Registry, int64, error) {
	var results []*image.Registry
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ImageRegistryDao) Create(ctx context.Context, in *image.Registry) (*image.Registry, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *ImageRegistryDao) Update(ctx context.Context, in *image.Registry) (*image.Registry, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
