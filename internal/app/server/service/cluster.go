package service

import (
	"context"
	"net/http"
	"sync"

	"github.com/kiaedev/kiae/api/cluster"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClusterService struct {
	clientStore sync.Map

	daoCluster *dao.ClusterDao
}

func NewClusterService(daoCluster *dao.ClusterDao) *ClusterService {
	return &ClusterService{daoCluster: daoCluster}
}

func (s *ClusterService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clusterId := r.Header.Get("X-Cluster-Id")
		if clusterId == "" {
			clusterId = r.URL.Query().Get("cluster-id")
		}
		if clusterId == "" {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		// ctx, err := s.WithKubeClientsCtx(ctx, clusterId)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *ClusterService) List(ctx context.Context, in *cluster.ListRequest) (*cluster.ListResponse, error) {
	results, total, err := s.daoCluster.List(ctx, bson.M{})
	return &cluster.ListResponse{Items: results, Total: total}, err
}

func (s *ClusterService) Create(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	eg, err := s.daoCluster.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return eg, nil
}

func (s *ClusterService) Update(ctx context.Context, in *cluster.UpdateRequest) (*cluster.Cluster, error) {
	return s.daoCluster.Update(ctx, in.Payload)
}

func (s *ClusterService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	_, err := s.daoCluster.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.daoCluster.Delete(ctx, in.Id); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
