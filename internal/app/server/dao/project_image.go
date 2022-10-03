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

type ProjectImageDao struct {
	*Dao
}

func NewProjectImageDao(db *mongo.Database) *ProjectImageDao {
	return &ProjectImageDao{
		Dao: NewDao(db.Collection("project-images")),
	}
}

func (p *ProjectImageDao) Get(ctx context.Context, id string) (*image.Image, error) {
	var rt image.Image
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ProjectImageDao) GetByName(ctx context.Context, name string) (*image.Image, error) {
	var rt image.Image
	if err := p.collection.FindOne(ctx, bson.M{"name": name}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ProjectImageDao) List(ctx context.Context, m bson.M) ([]*image.Image, int64, error) {
	var results []*image.Image
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ProjectImageDao) Create(ctx context.Context, in *image.Image) (*image.Image, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *ProjectImageDao) Update(ctx context.Context, in *image.Image) (*image.Image, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
