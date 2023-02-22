package log_test

import (
	"testing"
	"time"

	"go.arcalot.io/log/v2"
)

func TestMessage(t *testing.T) {
	timestamp := "2006-01-02T15:04:05Z"
	messageTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t.Fatal(err)
	}
	m := log.Message{
		Timestamp: messageTime,
		Level:     log.LevelError,
		Labels:    map[string]string{"source": "test"},
		Message:   "Hello world!",
	}
	if m.String() != "2006-01-02T15:04:05Z\terror\tsource=test\tHello world!" {
		t.Fatalf("Incorrect message string: %s", m.String())
	}
}
