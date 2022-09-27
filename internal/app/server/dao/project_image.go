package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/project"
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

func (p *ProjectImageDao) Get(ctx context.Context, id string) (*project.Image, error) {
	var rt project.Image
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ProjectImageDao) List(ctx context.Context, m bson.M) ([]*project.Image, int64, error) {
	var results []*project.Image
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ProjectImageDao) Create(ctx context.Context, in *project.Image) (*project.Image, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *ProjectImageDao) Update(ctx context.Context, in *project.Image) (*project.Image, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
