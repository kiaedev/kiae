package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/deployment"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DeploymentDao struct {
	*Dao
}

func NewDeploymentDao(db *mongo.Database) *DeploymentDao {
	return &DeploymentDao{
		Dao: NewDao(db.Collection("deployment")),
	}
}

func (p *DeploymentDao) Get(ctx context.Context, id string) (*deployment.Deployment, error) {
	var rt deployment.Deployment
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *DeploymentDao) GetByName(ctx context.Context, name string) (*deployment.Deployment, error) {
	var rt deployment.Deployment
	if err := p.collection.FindOne(ctx, bson.M{"name": name}).Decode(&rt); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *DeploymentDao) List(ctx context.Context, m bson.M) ([]*deployment.Deployment, int64, error) {
	var results []*deployment.Deployment
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *DeploymentDao) Create(ctx context.Context, in *deployment.Deployment) (*deployment.Deployment, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *DeploymentDao) Update(ctx context.Context, in *deployment.Deployment) (*deployment.Deployment, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
