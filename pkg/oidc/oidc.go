package oidc

type Config struct {
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
