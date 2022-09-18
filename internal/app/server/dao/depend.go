package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/depend"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DependDao struct {
	*Dao
}

func NewDependDao(db *mongo.Database) *DependDao {
	return &DependDao{
		Dao: NewDao(db.Collection("depends")),
	}
}

func (p *DependDao) Get(ctx context.Context, id string) (*depend.Depend, error) {
	var proj depend.Depend
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *DependDao) List(ctx context.Context, m bson.M) ([]*depend.Depend, int64, error) {
	var results []*depend.Depend
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *DependDao) Create(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *DependDao) Update(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
