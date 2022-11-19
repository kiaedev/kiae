package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/egress"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EgressDao struct {
	*Dao
}

func NewEgressDao(db *mongo.Database) *EgressDao {
	return &EgressDao{
		Dao: NewDao(db.Collection("dependent")),
	}
}

func (p *EgressDao) Get(ctx context.Context, id string) (*egress.Egress, error) {
	var proj egress.Egress
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *EgressDao) List(ctx context.Context, m bson.M) ([]*egress.Egress, int64, error) {
	var results []*egress.Egress
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *EgressDao) Create(ctx context.Context, in *egress.Egress) (*egress.Egress, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *EgressDao) Update(ctx context.Context, in *egress.Egress) (*egress.Egress, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
