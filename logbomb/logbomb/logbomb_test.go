package logbomb

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestNewWithInvalidLogWriterType(t *testing.T) {
	const logWriterType = "bogus"
	os.Setenv("LOG_WRITER_TYPE", logWriterType)
	_, err := NewLogBomb()
	if err == nil {
		t.Fatalf("Did not receive an error message")
	}
	unrecognizedErr, ok := err.(errUnrecognizedLogWriterType)
	if !ok {
		t.Fatalf("Expected an errUnrecognizedLogWriterType, received %s", reflect.TypeOf(err).String())
	}
	if unrecognizedErr.logWriterType != logWriterType {
		t.Fatalf("Got an errUnrecognizedLogWriterType, but expected logWriterType %s, got %s", logWriterType, unrecognizedErr.logWriterType)
	}
}

func TestNew(t *testing.T) {
	os.Setenv("LOG_WRITER_TYPE", "nsq")
	lb, err := NewLogBomb()
	if err != nil {
		t.Fatal(err)
	}
	logWriterType, ok := lb.logWriter.(*nsqLogWriter)
	if !ok {
		t.Fatalf("Expected a *logbomb.nsqLogWriter, got %s", reflect.TypeOf(logWriterType).String())
	}
}

func TestDetonate(t *testing.T) {
	lb, err := NewLogBomb()
	if err != nil {
		t.Fatal(err)
	}
	// Swap in a mock logWriter
	mlw := &mockLogWriter{}
	lb.logWriter = mlw
	if err := lb.Detonate(); err != nil {
		t.Fatal(err)
	}
	if !mlw.written {
		t.Fatal("Expected mockLogWriter's write() function to have been called, but it was not.")
	}
}

type mockLogWriter struct {
	written bool
	mutex   sync.Mutex
}

func (lw *mockLogWriter) write(message string) error {
	lw.mutex.Lock()
	defer lw.mutex.Unlock()
	lw.written = true
	return nil
}
