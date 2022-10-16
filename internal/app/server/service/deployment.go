package service

import (
	"context"

	"github.com/kiaedev/kiae/api/deployment"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeploymentService struct {
	deploymentDao *dao.DeploymentDao

	imageSvc *ProjectImageSvc
	appSvc   *AppService
}

func NewDeploymentService(deploymentDao *dao.DeploymentDao, imageSvc *ProjectImageSvc, appSvc *AppService) *DeploymentService {
	return &DeploymentService{deploymentDao: deploymentDao, imageSvc: imageSvc, appSvc: appSvc}
}

func (s *DeploymentService) List(ctx context.Context, in *deployment.DeploymentListRequest) (*deployment.DeploymentListResponse, error) {
	results, total, err := s.deploymentDao.List(ctx, bson.M{"pid": in.Pid})
	return &deployment.DeploymentListResponse{Items: results, Total: total}, err
}

func (s *DeploymentService) Create(ctx context.Context, in *deployment.Deployment) (*deployment.Deployment, error) {
	img, err := s.imageSvc.daoProjImg.Get(ctx, in.ImageId)
	if err != nil {
		return nil, err
	}

	ap, err := s.appSvc.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	ap.Image = img.Image
	if err := s.appSvc.updateAppComponent(ctx, ap); err != nil {
		return nil, err
	}

	in.ImageUrl = img.Image
	in.CommitId = img.CommitId
	in.CommitMsg = img.CommitMsg
	in.CommitAuthor = img.CommitAuthor
	return s.deploymentDao.Create(ctx, in)
}

func (s *DeploymentService) Update(ctx context.Context, in *deployment.Deployment) (*deployment.Deployment, error) {
	return s.deploymentDao.Update(ctx, in)
}

func (s *DeploymentService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.deploymentDao.Delete(ctx, in.Id)
}
