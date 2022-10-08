//go:build wireinject
// +build wireinject

package server

import (
	"fmt"

	"github.com/google/wire"
	"github.com/kiaedev/kiae/api/graph"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	"github.com/kiaedev/kiae/internal/pkg/kcs"
	"github.com/kiaedev/kiae/pkg/loki"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/rest"
)

func dbConstructor() (*mongo.Database, error) {
	dbClient, err := mongoutil.New(viper.GetString("dsn"))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	return dbClient.DB.Database("kiae"), nil
}

func lokiConstructor() (*loki.Client, error) {
	return loki.NewLoki("http://localhost:3100"), nil
}

func buildInjectors(config *rest.Config) (*Server, error) {
	wire.Build(
		dbConstructor,
		lokiConstructor,
		kcs.ProviderSet,
		dao.ProviderSet,
		service.ProviderSet,
		watch.NewWatcher,
		wire.Struct(new(graph.Resolver), "*"),
		wire.Struct(new(Server), "*"),
	)

	return &Server{}, nil
}
