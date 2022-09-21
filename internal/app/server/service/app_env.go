package service

import (
	"context"

	"github.com/kiaedev/kiae/api/app"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AppService) EnvCreate(ctx context.Context, in *app.AppEnv) (*app.Environment, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}

	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "application not found by the id %v", ap.Id)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	for _, environment := range ap.Environments {
		if environment.Name == in.Payload.Name {
			return nil, status.Errorf(codes.AlreadyExists, "environment already exist")
		}
	}

	in.Payload.Type = app.Environment_USER
	in.Payload.CreatedAt = timestamppb.Now()
	in.Payload.UpdatedAt = timestamppb.Now()
	ap.Environments = append(ap.Environments, in.Payload)
	_, err = s.daoApp.Update(ctx, ap)
	return in.Payload, err
}

func (s *AppService) EnvUpdate(ctx context.Context, in *app.AppEnv) (*app.Environment, error) {
	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	for idx, env := range ap.Environments {
		if env.Name == in.Payload.Name && env.Type > app.Environment_SYSTEM {
			ap.Environments[idx].Value = in.Payload.Value
			ap.Environments[idx].UpdatedAt = timestamppb.Now()
		}
	}

	_, err = s.daoApp.Update(ctx, ap)
	return in.Payload, err
}

func (s *AppService) EnvDelete(ctx context.Context, in *app.AppEnv) (*emptypb.Empty, error) {
	ap, err := s.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	for idx, environment := range ap.Environments {
		if environment.Name == in.Payload.Name {
			ap.Environments = append(ap.Environments[:idx], ap.Environments[idx+1:]...)
		}
	}

	_, err = s.daoApp.Update(ctx, ap)
	return &emptypb.Empty{}, err
}
