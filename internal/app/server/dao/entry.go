package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/entry"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EntryDao struct {
	*Dao
}

func NewEntryDao(db *mongo.Database) *EntryDao {
	return &EntryDao{
		Dao: NewDao(db.Collection("entries")),
	}
}

func (p *EntryDao) Get(ctx context.Context, id string) (*entry.Entry, error) {
	var proj entry.Entry
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *EntryDao) List(ctx context.Context, m bson.M) ([]*entry.Entry, int64, error) {
	var results []*entry.Entry
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *EntryDao) Create(ctx context.Context, in *entry.Entry) (*entry.Entry, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *EntryDao) Update(ctx context.Context, in *entry.Entry) (*entry.Entry, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
