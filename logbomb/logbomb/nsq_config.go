package logbomb

import (
	"github.com/kelseyhightower/envconfig"
)

type nsqConfig struct {
	Host  string `envconfig:"DEIS_NSQD_SERVICE_HOST" default:""`
	Port  int    `envconfig:"DEIS_NSQD_SERVICE_PORT_TRANSPORT" default:"4150"`
	Topic string `envconfig:"NSQ_TOPIC" default:"logs"`
}

func parseNSQConfig() (*nsqConfig, error) {
	ret := new(nsqConfig)
	if err := envconfig.Process(appName, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
