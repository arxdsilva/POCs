package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PrivateKeyFile string `envconfig:"APP_JWT_PRIVATE_KEY_FILE" default:"/data/default.pem"`
	// PrivateKeyFile string `envconfig:"APP_JWT_PRIVATE_KEY_FILE" default:"default.pem"`
}

func FromEnv() (Config, error) {
	var config Config
	return config, envconfig.Process("APP", &config)
}
