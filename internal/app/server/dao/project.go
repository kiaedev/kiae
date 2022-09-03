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

type ProjectDao struct {
	*Dao
}

func NewProject(db *mongo.Database) *ProjectDao {
	return &ProjectDao{
		Dao: NewDao(db.Collection("projects")),
	}
}

func (p *ProjectDao) Get(ctx context.Context, id string) (*project.Project, error) {
	var proj project.Project
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := p.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *ProjectDao) List(ctx context.Context, m bson.M) ([]*project.Project, int64, error) {
	var results []*project.Project
	total, err := mongoutil.ListAndCount(ctx, p.collection, m, &results)
	return results, total, err
}

func (p *ProjectDao) Create(ctx context.Context, in *project.Project) (*project.Project, error) {
	in.CreatedAt = timestamppb.Now()
	in.UpdatedAt = timestamppb.Now()
	rt, err := p.collection.InsertOne(ctx, in)
	in.Id = rt.InsertedID.(primitive.ObjectID).Hex()
	return in, err
}

func (p *ProjectDao) Update(ctx context.Context, in *project.Project) (*project.Project, error) {
	oid, _ := primitive.ObjectIDFromHex(in.Id)
	in.UpdatedAt = timestamppb.Now()
	_, err := p.collection.UpdateByID(ctx, oid, in)
	return in, err
}
