// +build testnsq

package logbomb

import (
	"fmt"
	"testing"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

func TestWrite(t *testing.T) {
	message := "foobar"
	lw, err := newNSQLogWriter()
	if err != nil {
		t.Fatal(err)
	}
	nlw, ok := lw.(*nsqLogWriter)
	if !ok {
		t.Fatal(err)
	}
	nsqConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(nlw.config.Topic, "test", nsqConfig)
	if err != nil {
		t.Fatal(err)
	}
	foundCh := make(chan bool, 1)
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		if string(msg.Body) == message {
			fmt.Println(msg.Body)
			foundCh <- true
		}
		return nil
	}))
	if err := consumer.ConnectToNSQD(fmt.Sprintf("%s:%d", nlw.config.Host, nlw.config.Port)); err != nil {
		t.Fatal(err)
	}
	err = nlw.write(message)
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-foundCh:
		return
	case <-time.After(time.Second * 10):
		t.Fatal("Timed out waiting for message.")
	}
}
