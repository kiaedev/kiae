package service

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/route"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RouteService struct {
	route.UnimplementedRouteServiceServer

	appSvc   *AppService
	daoRoute *dao.RouteDao
}

func NewRouteService(appSvc *AppService, daoRoute *dao.RouteDao) *RouteService {
	return &RouteService{appSvc: appSvc, daoRoute: daoRoute}
}

func (s *RouteService) List(ctx context.Context, in *route.ListRequest) (*route.ListResponse, error) {
	results, total, err := s.daoRoute.List(ctx, bson.M{"appid": in.Appid})
	return &route.ListResponse{Items: results, Total: total}, err
}

func (s *RouteService) Create(ctx context.Context, in *route.Route) (*route.Route, error) {
	return s.daoRoute.Create(ctx, in)
}

func (s *RouteService) Update(ctx context.Context, in *route.UpdateRequest) (*route.Route, error) {
	existedRoute, err := s.daoRoute.Get(ctx, in.Payload.Id)
	if err != nil {
		return nil, err
	}

	if in.UpdateMask == nil {
		existedRoute = in.Payload
	} else {
		s.handlePatch(in, existedRoute)
	}

	fRoute, err := s.daoRoute.Update(ctx, existedRoute)
	if err != nil {
		return nil, err
	}

	return fRoute, s.appSvc.updateAppComponentById(ctx, fRoute.Appid)
}

func (s *RouteService) handlePatch(in *route.UpdateRequest, existedRoute *route.Route) {
	payload := in.Payload
	for _, path := range in.GetUpdateMask().Paths {
		if path == "status" {
			existedRoute.Status = payload.Status
		}
	}
}

func (s *RouteService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoRoute.Delete(ctx, in.Id)
}
