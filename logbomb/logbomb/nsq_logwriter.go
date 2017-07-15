package logbomb

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

type nsqLogWriter struct {
	config   *nsqConfig
	producer *nsq.Producer
}

func newNSQLogWriter() (logWriter, error) {
	cfg, err := parseNSQConfig()
	if err != nil {
		return nil, err
	}
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nsqConfig)
	if err != nil {
		return nil, err
	}
	return &nsqLogWriter{
		config:   cfg,
		producer: producer,
	}, nil
}

func (lw *nsqLogWriter) write(message string) error {
	if err := lw.producer.Publish(lw.config.Topic, []byte(message)); err != nil {
		return err
	}
	return nil
}
