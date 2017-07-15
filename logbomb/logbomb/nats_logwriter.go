package logbomb

import (
	"github.com/nats-io/go-nats"
)

type natsLogWriter struct {
	config *natsConfig
	conn   *nats.Conn
}

func newNATSLogWriter() (*natsLogWriter, error) {

	cfg, err := parseNATSConfig()
	if err != nil {
		return nil, err
	}

	conn, err := nats.Connect(cfg.URL)
	if err != nil {
		return nil, err
	}

	lw := &natsLogWriter{conn: conn, config: cfg}
	return lw, nil
}

func (lw *natsLogWriter) write(message string) error {

	if err := lw.conn.Publish(lw.config.Topic, []byte(message)); err != nil {
		return err
	}
	return nil
}
