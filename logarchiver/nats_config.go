package main

import (
	"github.com/kelseyhightower/envconfig"
)

type natsConfig struct {
	URL   string `envconfig:"NATS_URL"`
	Topic string `envconfig:"NATS_TOPIC" default:"logs"`
}

func parseNATSConfig() (*natsConfig, error) {
	ret := new(natsConfig)
	if err := envconfig.Process(appName, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
