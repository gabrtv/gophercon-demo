package logbomb

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	appName = "logbomb"
)

type config struct {
	GoRoutines           int    `envconfig:"GO_ROUTINES" default:"100"`
	MessagesPerGoRoutine int    `envconfig:"MESSAGES_PER_GO_ROUTINE" default:"100"`
	LogWriterType        string `envconfig:"LOG_WRITER_TYPE" default:"nats"`
	MinMessageWords      int    `envconfig:"MIN_MESSAGE_WORDS" default:"5"`
	MaxMessageWords      int    `envconfig:"MAX_MESSAGE_WORDS" default:"50"`
}

func parseConfig() (*config, error) {
	ret := new(config)
	if err := envconfig.Process(appName, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
