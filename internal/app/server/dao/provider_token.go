package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProviderTokenDao struct {
	*Dao
}

func NewProviderTokenDao(db *mongo.Database) *ProviderTokenDao {
	return &ProviderTokenDao{
		Dao: NewDao(db.Collection("provider-tokens")),
	}
}

func (p *ProviderTokenDao) GetByProvider(ctx context.Context, name string) (*provider.Token, error) {
	var rt provider.Token
	if err := p.collection.FindOne(ctx, bson.M{"provider": name}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *ProviderTokenDao) Create(ctx context.Context, in *provider.Token) (*provider.Token, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	if err == nil {
		in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	}
	return in, err
}

func (p *ProviderTokenDao) Upsert(ctx context.Context, in *provider.Token) (*provider.Token, error) {
	in.Id = "" // clean the immutable field
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"provider": in.Provider}, bson.M{"$set": in}, options.Update().SetUpsert(true))
	return in, err
}

func (p *ProviderTokenDao) Update(ctx context.Context, in *provider.Token) (*provider.Token, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	in.Id = oid.Hex()
	return in, err
}
