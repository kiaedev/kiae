package dao

import (
	"context"

	"github.com/kiaedev/kiae/api/cluster"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ClusterDao struct {
	*Dao
}

func NewClusterDao(db *mongo.Database) *ClusterDao {
	return &ClusterDao{
		Dao: NewDao(db.Collection("cluster")),
	}
}

func (p *ClusterDao) Get(ctx context.Context, id string) (*cluster.Cluster, error) {
	var proj cluster.Cluster
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *ClusterDao) List(ctx context.Context, m bson.M) ([]*cluster.Cluster, int64, error) {
	var results []*cluster.Cluster
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ClusterDao) Create(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *ClusterDao) Update(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.Id = "" // clean the immutable field
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": in})
	return in, err
}
