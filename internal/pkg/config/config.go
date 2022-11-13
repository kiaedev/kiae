package config

import (
	"github.com/kiaedev/kiae/pkg/oauth2"
	"github.com/spf13/viper"
)

type Config struct {
	Oidc *oauth2.OidcConfig `yaml:"oidc"`
}

func New() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
