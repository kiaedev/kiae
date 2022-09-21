package service

import (
	"context"

	"github.com/kiaedev/kiae/api/app"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AppService) CfgCreate(ctx context.Context, in *app.AppCfg) (*app.Configuration, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}

	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "application not found by the id %v", ap.Id)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	for _, cfg := range ap.Configs {
		if cfg.Filename == in.Payload.Filename && cfg.MountPath == in.Payload.MountPath {
			return nil, status.Errorf(codes.AlreadyExists, "config already exist")
		}
	}

	in.Payload.Name = strutil.RandomText(5)
	in.Payload.CreatedAt = timestamppb.Now()
	in.Payload.UpdatedAt = timestamppb.Now()
	ap.Configs = append(ap.Configs, in.Payload)
	_, err = s.daoApp.Update(ctx, ap)
	return in.Payload, err
}

func (s *AppService) CfgUpdate(ctx context.Context, in *app.AppCfg) (*app.Configuration, error) {
	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	for idx, cfg := range ap.Configs {
		if cfg.Name == in.Payload.Name {
			ap.Configs[idx].Filename = in.Payload.Filename
			ap.Configs[idx].MountPath = in.Payload.MountPath
			ap.Configs[idx].Content = in.Payload.Content
			ap.Configs[idx].UpdatedAt = in.Payload.UpdatedAt
		}
	}

	_, err = s.daoApp.Update(ctx, ap)
	return in.Payload, err
}

func (s *AppService) CfgDelete(ctx context.Context, in *app.AppCfg) (*emptypb.Empty, error) {
	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	for idx, cfg := range ap.Configs {
		if cfg.Name == in.Payload.Name {
			ap.Configs = append(ap.Configs[:idx], ap.Configs[idx+1:]...)
		}
	}

	_, err = s.daoApp.Update(ctx, ap)
	return &emptypb.Empty{}, err
}
