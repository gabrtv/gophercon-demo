package main

import (
	"log"
	"time"

	"sync"

	"bytes"

	"bufio"

	"github.com/nats-io/nats"
)

const (
	appName = "logarchiver"
)

func main() {
	log.Println("starting log archiver...")

	natsConfig, err := parseNATSConfig()
	if err != nil {
		log.Fatalf("failed to parse NATS configuration: %v\n", err)
	}

	minioConfig, err := parseMinioConfig()
	if err != nil {
		log.Fatalf("failed to parse Minio configuration: %v\n", err)
	}

	ch := make(chan *nats.Msg, 64)

	var wg sync.WaitGroup

	wg.Add(2)
	go recv(natsConfig, ch, &wg)
	go send(minioConfig, ch, &wg)
	wg.Wait()

}

func recv(cfg *natsConfig, ch chan *nats.Msg, wg *sync.WaitGroup) {
	defer wg.Done()

	c, err := nats.Connect(cfg.URL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v\n", err)
	}
	defer c.Close()

	_, err = c.ChanSubscribe(cfg.Topic, ch)

	for c.IsConnected() && !c.IsClosed() {
	}
}

func send(cfg *minioConfig, ch chan *nats.Msg, wg *sync.WaitGroup) {
	defer wg.Done()

	mc, err := newMinioClient(cfg)
	if err != nil {
		log.Fatalf("failed to connect to Minio: %v\n", err)
	}

	err = createBucket(cfg, mc)
	if err != nil {
		log.Fatalf("failed to create bucket: %v\n", err)
	}

	var b bytes.Buffer
	r := bufio.NewReader(&b)
	tick := time.Tick(10 * time.Second)

	for {
		select {
		case msg := <-ch:
			b.Write(msg.Data)
		case t := <-tick:
			if b.Len() > 0 {
				uploadFile(cfg, mc, r, t)
			}
		}
	}
}
