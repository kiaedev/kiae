package service

import (
	"context"
	"strings"

	"github.com/kiaedev/kiae/api/depend"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DependService struct {
	depend.UnimplementedDependServiceServer

	daoApp    *dao.AppDao
	daoDepend *dao.DependDao
	// daoMiddleware *dao.Middleware
	daoMwInstance *dao.MiddlewareInstance
}

func NewDependService(cs *Service) *DependService {
	return &DependService{
		daoApp:        dao.NewApp(cs.DB),
		daoDepend:     dao.NewDependDao(cs.DB),
		daoMwInstance: dao.NewMiddlewareInstanceDao(cs.DB),
	}
}

func (s *DependService) List(ctx context.Context, in *depend.ListRequest) (*depend.ListResponse, error) {
	results, total, err := s.daoDepend.List(ctx, bson.M{"appid": in.Appid})
	return &depend.ListResponse{Items: results, Total: total}, err
}

func (s *DependService) Create(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	// kApp, err := s.daoApp.Get(ctx, in.Appid)
	// if err != nil {
	// 	return nil, err
	// }

	in.Name = strings.ToLower(strutil.RandomText(8))
	if in.Type == depend.Depend_MIDDLEWARE && in.MInstance != "" {
		in.Status = depend.Depend_BOUND
	}

	return s.daoDepend.Create(ctx, in)
}

func (s *DependService) Update(ctx context.Context, in *depend.Depend) (*depend.Depend, error) {
	return s.daoDepend.Update(ctx, in)
}

func (s *DependService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoDepend.Delete(ctx, in.Id)
}
