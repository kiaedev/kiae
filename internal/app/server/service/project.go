package service

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer

	daoProj *dao.ProjectDao
}

func NewProjectService(s *Service) *ProjectService {
	return &ProjectService{
		daoProj: dao.NewProject(s.DB),
	}
}

func (p *ProjectService) List(ctx context.Context, in *project.ListRequest) (*project.ListResponse, error) {
	results, total, err := p.daoProj.List(ctx, bson.M{})
	return &project.ListResponse{Items: results, Total: total}, err
}

func (p *ProjectService) Create(ctx context.Context, in *project.Project) (*project.Project, error) {
	setDefaultProjectProperties(in)
	return p.daoProj.Create(ctx, in)
}

func (p *ProjectService) Update(ctx context.Context, in *project.Project) (*project.Project, error) {
	return p.daoProj.Update(ctx, in)
}

func (p *ProjectService) Read(ctx context.Context, in *kiae.IdRequest) (*project.Project, error) {
	return p.daoProj.Get(ctx, in.Id)
}

func setDefaultProjectProperties(in *project.Project) {
	imageRegistry := "saltbo/"
	in.ImageRepo = imageRegistry + in.Name

	// todo 从配置中获取镜像仓库地址
	// in.Images = []*project.Image{
	// 	{
	// 		Name:      in.Name,
	// 		Image:     imageRegistry + in.Name,
	// 		Latest:    "latest",
	// 		CreatedAt: timestamppb.Now(),
	// 		UpdatedAt: timestamppb.Now(),
	// 	},
	// }
}
