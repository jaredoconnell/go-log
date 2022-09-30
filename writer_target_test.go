package log //nolint:testpackage

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type testTarget struct {
	data []byte
}

func (t *testTarget) Log(args ...interface{}) {
	t.data = append(t.data, []byte(fmt.Sprint(args...)+"\n")...)
}

func (t *testTarget) Logf(message string, args ...interface{}) {
	t.data = append(t.data, []byte(fmt.Sprintf(message, args...)+"\n")...)
}

func TestTargetWriter_Write(t *testing.T) {
	target := &testTarget{}
	logger := newTargetWriter(target.Log)
	if err := logger.Write(Message{
		Timestamp: time.Now(),
		Level:     LevelError,
		Labels:    nil,
		Message:   "Hello world!",
	}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(target.data), "Hello world!") {
		t.Fatalf("Incorrect log output: %s", target.data)
	}
}
