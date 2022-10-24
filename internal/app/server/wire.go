//go:build wireinject
// +build wireinject

package server

import (
	"fmt"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/kiaedev/kiae/api/graph"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	"github.com/kiaedev/kiae/internal/pkg/config"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/kiaedev/kiae/pkg/loki"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/kiaedev/kiae/pkg/oauth2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/rest"
)

func dbConstructor() (*mongo.Database, error) {
	dbClient, err := mongoutil.New(viper.GetString("mongodb.dsn"))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	return dbClient.DB.Database("kiae"), nil
}

func lokiConstructor() (*loki.Client, error) {
	return loki.NewLoki(viper.GetString("loki.endpoint")), nil
}

func buildInjectors(kubeconfig *rest.Config) (*Server, error) {
	wire.Build(
		config.New,
		dbConstructor,
		lokiConstructor,
		klient.ProviderSet,
		wire.FieldsOf(new(*config.Config), "OIDC"),
		oauth2.NewOIDC,
		mux.NewRouter,
		dao.ProviderSet,
		service.ProviderSet,
		watch.NewWatcher,
		watch.NewMultiClusterInformers,
		wire.Struct(new(graph.Resolver), "*"),
		wire.Struct(new(Server), "*"),
	)

	return &Server{}, nil
}
