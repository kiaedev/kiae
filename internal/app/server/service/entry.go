package service

import (
	"context"

	"github.com/kiaedev/kiae/api/entry"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EntryService struct {
	entry.UnimplementedEntryServiceServer

	appSvc   *AppService
	daoEntry *dao.EntryDao
}

func NewEntryService(appSvc *AppService, daoEntry *dao.EntryDao) *EntryService {
	return &EntryService{appSvc: appSvc, daoEntry: daoEntry}
}

func (s *EntryService) List(ctx context.Context, in *entry.ListRequest) (*entry.ListResponse, error) {
	results, total, err := s.daoEntry.List(ctx, bson.M{"appid": in.Appid})
	return &entry.ListResponse{Items: results, Total: total}, err
}

func (s *EntryService) Create(ctx context.Context, in *entry.Entry) (*entry.Entry, error) {
	return s.daoEntry.Create(ctx, in)
}

func (s *EntryService) Update(ctx context.Context, in *entry.UpdateRequest) (*entry.Entry, error) {
	existedEntry, err := s.daoEntry.Get(ctx, in.Payload.Id)
	if err != nil {
		return nil, err
	}

	if in.UpdateMask == nil {
		existedEntry = in.Payload
	} else {
		s.handlePatch(in, existedEntry)
	}

	fEntry, err := s.daoEntry.Update(ctx, existedEntry)
	if err != nil {
		return nil, err
	}

	return fEntry, s.appSvc.updateAppComponentById(ctx, fEntry.Appid)
}

func (s *EntryService) handlePatch(in *entry.UpdateRequest, existedEntry *entry.Entry) {
	payload := in.Payload
	for _, path := range in.GetUpdateMask().Paths {
		if path == "status" {
			existedEntry.Status = payload.Status
		}
	}
}

func (s *EntryService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoEntry.Delete(ctx, in.Id)
}
