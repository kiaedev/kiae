package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProviderDao struct {
	*Dao
}

func NewProviderDao(db *mongo.Database) *ProviderDao {
	return &ProviderDao{
		Dao: NewDao(db.Collection("providers")),
	}
}

func (p *ProviderDao) Get(ctx context.Context, id string) (*provider.Provider, error) {
	var proj provider.Provider
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *ProviderDao) GetByName(ctx context.Context, name string) (*provider.Provider, error) {
	var proj provider.Provider
	if err := p.collection.FindOne(ctx, bson.M{"name": name}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *ProviderDao) List(ctx context.Context, m bson.M) ([]*provider.Provider, int64, error) {
	var results []*provider.Provider
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ProviderDao) Create(ctx context.Context, in *provider.Provider) (*provider.Provider, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	if err == nil {
		in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	}
	return in, err
}

func (p *ProviderDao) Update(ctx context.Context, in *provider.Provider) (*provider.Provider, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	in.Id = oid.Hex()
	return in, err
}
