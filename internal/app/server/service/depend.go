package service

import (
	"context"
	"strings"

	"github.com/kiaedev/kiae/api/depend"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DependService struct {
	depend.UnimplementedDependServiceServer

	appSvc    *AppService
	daoDepend *dao.DependDao
	// daoMiddleware *dao.Middleware
	daoMwInstance *dao.MiddlewareInstance
}

func NewDependService(cs *Service) *DependService {
	return &DependService{
		appSvc:        NewAppService(cs),
		daoDepend:     dao.NewDependDao(cs.DB),
		daoMwInstance: dao.NewMiddlewareInstanceDao(cs.DB),
	}
}

func (s *DependService) List(ctx context.Context, in *depend.ListRequest) (*depend.ListResponse, error) {
	results, total, err := s.daoDepend.List(ctx, bson.M{"appid": in.Appid})
	return &depend.ListResponse{Items: results, Total: total}, err
}

func (s *DependService) Create(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	in.Name = strings.ToLower(strutil.RandomText(8))
	if in.Type == depend.Depend_MIDDLEWARE && in.MInstance != "" {
		in.Status = depend.Depend_BOUND
	}

	if err := s.appSvc.addComponent(ctx, in.Appid, components.MwConstructor(in.MType, in.MInstance, in.Name)); err != nil {
		return nil, err
	}

	return s.daoDepend.Create(ctx, in)
}

func (s *DependService) Update(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	return s.daoDepend.Update(ctx, in)
}

func (s *DependService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	dpb, err := s.daoDepend.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.appSvc.removeComponent(ctx, dpb.Appid, components.MwConstructor(dpb.MType, dpb.MInstance, dpb.Name)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoDepend.Delete(ctx, in.Id)
}
