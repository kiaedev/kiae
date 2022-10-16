package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/builder"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BuilderDao struct {
	*Dao
}

func NewBuilderDao(db *mongo.Database) *BuilderDao {
	return &BuilderDao{
		Dao: NewDao(db.Collection("builders")),
	}
}

func (p *BuilderDao) Get(ctx context.Context, id string) (*builder.Builder, error) {
	var rt builder.Builder
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *BuilderDao) GetByName(ctx context.Context, name string) (*builder.Builder, error) {
	var rt builder.Builder
	if err := p.collection.FindOne(ctx, bson.M{"name": name}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *BuilderDao) List(ctx context.Context, m bson.M) ([]*builder.Builder, int64, error) {
	var results []*builder.Builder
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *BuilderDao) Create(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *BuilderDao) Update(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
