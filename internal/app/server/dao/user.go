package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/user"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserDao struct {
	*Dao
}

func NewUserDao(db *mongo.Database) *UserDao {
	return &UserDao{
		Dao: NewDao(db.Collection("user")),
	}
}

func (p *UserDao) Get(ctx context.Context, id string) (*user.User, error) {
	var proj user.User
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *UserDao) GetByOuterId(ctx context.Context, subject string) (*user.User, error) {
	var u user.User
	if err := p.collection.FindOne(ctx, bson.M{"outer_id": subject}).Decode(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *UserDao) List(ctx context.Context, m bson.M) ([]*user.User, int64, error) {
	var results []*user.User
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *UserDao) Create(ctx context.Context, in *user.User) (*user.User, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *UserDao) Update(ctx context.Context, in *user.User) (*user.User, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
