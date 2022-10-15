package service

import (
	"context"

	"github.com/kiaedev/kiae/api/cluster"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ClusterService struct {
	daoCluster *dao.ClusterDao

	wci *watch.MultiClusterInformers
}

func NewClusterService(daoCluster *dao.ClusterDao, wci *watch.MultiClusterInformers) *ClusterService {
	wci.SetupStockClusterFetcher(func(ctx context.Context) []*cluster.Cluster {
		clusters, _, _ := daoCluster.List(ctx, bson.M{})
		return wrapLocalCluster(clusters)
	})
	return &ClusterService{daoCluster: daoCluster, wci: wci}
}

func (s *ClusterService) List(ctx context.Context, in *cluster.ListRequest) (*cluster.ListResponse, error) {
	results, total, err := s.daoCluster.List(ctx, bson.M{})
	return &cluster.ListResponse{Items: wrapLocalCluster(results), Total: total}, err
}

func (s *ClusterService) Create(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	// TODO: create a ClusterGateway CR

	eg, err := s.daoCluster.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	s.wci.ClusterEvent(watch.ClusterEventAddon, eg)
	return eg, nil
}

func (s *ClusterService) Update(ctx context.Context, in *cluster.UpdateRequest) (*cluster.Cluster, error) {
	return s.daoCluster.Update(ctx, in.Payload)
}

func (s *ClusterService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	eg, err := s.daoCluster.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// TODO: remove the ClusterGateway

	if err := s.daoCluster.Delete(ctx, in.Id); err != nil {
		return &emptypb.Empty{}, err
	}

	s.wci.ClusterEvent(watch.ClusterEventRemoved, eg)
	return &emptypb.Empty{}, nil
}

func wrapLocalCluster(clusters []*cluster.Cluster) []*cluster.Cluster {
	return append(clusters, &cluster.Cluster{
		Name:      "local",
		Intro:     "local control plane cluster",
		CreatedAt: timestamppb.Now(), // todo use the kiae installed time
	})
}
