package service

import (
	"context"
	"fmt"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProjectService struct {
	daoProj *dao.ProjectDao
}

func NewProjectService(daoProj *dao.ProjectDao) *ProjectService {
	return &ProjectService{daoProj: daoProj}
}

func (p *ProjectService) List(ctx context.Context, in *project.ListRequest) (*project.ListResponse, error) {
	query := bson.M{"owner_uid": MustGetUserid(ctx)}
	// todo support GroupMember display

	results, total, err := p.daoProj.List(ctx, query)
	return &project.ListResponse{Items: results, Total: total}, err
}

func (p *ProjectService) Create(ctx context.Context, in *project.Project) (*project.Project, error) {
	in.OwnerUid = MustGetUserid(ctx)
	_, total, err := p.daoProj.List(ctx, bson.M{"owner_uid": in.GetOwnerUid(), "name": in.GetName()})
	if err == nil && total != 0 {
		return nil, fmt.Errorf("project %s already exists", in.GetName())
	} else if err != nil {
		return nil, err
	}

	return p.daoProj.Create(ctx, in)
}

func (p *ProjectService) Update(ctx context.Context, in *project.Project) (*project.Project, error) {
	return p.daoProj.Update(ctx, in)
}

func (p *ProjectService) Read(ctx context.Context, in *kiae.IdRequest) (*project.Project, error) {
	return p.daoProj.Get(ctx, in.Id)
}

func (p *ProjectService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.daoProj.Delete(ctx, in.Id)
}
