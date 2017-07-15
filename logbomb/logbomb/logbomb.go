package logbomb

import (
	"log"
	"sync"
)

// LogBomb places load directly on Deis Workflow's logging subsystem.
type LogBomb struct {
	config    *config
	logWriter logWriter
}

// NewLogBomb returns a new LogBomb.
func NewLogBomb() (*LogBomb, error) {
	cfg, err := parseConfig()
	if err != nil {
		return nil, err
	}
	var lw logWriter
	if cfg.LogWriterType == "nsq" {
		lw, err = newNSQLogWriter()
		if err != nil {
			return nil, err
		}
	} else if cfg.LogWriterType == "nats" {
		lw, err = newNATSLogWriter()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, newErrUnrecognizedLogWriterType(cfg.LogWriterType)
	}
	return &LogBomb{
		config:    cfg,
		logWriter: lw,
	}, nil
}

// Detonate places load directly on the logging subsystem.
func (lb *LogBomb) Detonate() error {
	var wg sync.WaitGroup
	for i := 0; i < lb.config.GoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				for j := 0; j < lb.config.MessagesPerGoRoutine; j++ {
					if err := lb.logWriter.write(lb.getMessage()); err != nil {
						log.Printf("Error writing log message: %s", err)
						log.Fatal("Shutting down. This is to avoid inadvertently producing MORE log messages while NSQ is already not responsing.")
					}
				}
			}
		}()
	}
	wg.Wait()
	return nil
}
