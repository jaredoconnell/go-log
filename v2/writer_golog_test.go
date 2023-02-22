package log_test

import (
	"bytes"
	goLog "log"
	"strings"
	"testing"
	"time"

	"go.arcalot.io/log/v2"
)

func TestGoLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	backingLogger := goLog.New(buf, "", 0)
	logger := log.NewGoLogWriter(backingLogger)
	err := logger.Write(log.Message{
		Timestamp: time.Now(),
		Level:     log.LevelError,
		Labels:    nil,
		Message:   "Hello world!",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "Hello world!\n") {
		t.Fatalf("failed to find written log message in output")
	}
}
