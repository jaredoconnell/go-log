package log_test

import (
	"bytes"
	"strings"
	"testing"

	"go.arcalot.io/log/v2"
)

//nolint:funlen
func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("stdout", func(t *testing.T) {
		t.Parallel()
		buf := &bytes.Buffer{}
		l := log.New(log.Config{
			Level:       log.LevelDebug,
			Destination: log.DestinationStdout,
			Stdout:      buf,
		})
		l.Debugf("Hello world!")
		data := buf.String()
		if !strings.Contains(data, "Hello world!") {
			t.Fatalf("Incorrect log output: %s", data)
		}
	})

	t.Run("testing", func(t *testing.T) {
		t.Parallel()
		buf := &bytes.Buffer{}
		testingT := testing.T{}
		l := log.New(log.Config{
			Level:       log.LevelDebug,
			Destination: log.DestinationTest,
			Stdout:      buf,
			T:           &testingT,
		})
		l.Debugf("Hello world!")
		if buf.Len() != 0 {
			t.Fatalf("Incorrect stdout content: %s", buf.String())
		}
		// There is no way to test logging via testing.T because the logging is happening in unexported functions.
	})

	t.Run("config-error", func(t *testing.T) {
		t.Parallel()
		testCases := map[string]log.Config{
			"missing-test": {
				Level:       log.LevelDebug,
				Destination: log.DestinationTest,
			},
			"invalid-level": {
				Level:       log.Level("invalid"),
				Destination: log.DestinationStdout,
			},
			"invalid-destination": {
				Level:       log.LevelDebug,
				Destination: log.Destination("invalid"),
			},
		}
		for name, tc := range testCases {
			testCase := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				var capturedError any
				func() {
					defer func() {
						capturedError = recover()
					}()
					log.New(testCase)
				}()
				if capturedError == nil {
					t.Fatalf("No error captured")
				}
			})
		}
	})
}

//nolint:dupl
func TestLogger_Debugf(t *testing.T) {
	t.Parallel()
	t.Run("unfiltered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelDebug, writer)
		logger.Debugf("Hello world!")
		if output := writer.String(); !strings.Contains(output, "Hello world!") || !strings.Contains(output, "debug") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})

	t.Run("filtered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelInfo, writer)
		logger.Debugf("Hello world!")
		if output := writer.String(); strings.Contains(output, "Hello world!") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})
}

//nolint:dupl
func TestLogger_Infof(t *testing.T) {
	t.Parallel()
	t.Run("unfiltered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelInfo, writer)
		logger.Infof("Hello world!")
		if output := writer.String(); !strings.Contains(output, "Hello world!") || !strings.Contains(output, "info") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})

	t.Run("filtered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelWarning, writer)
		logger.Infof("Hello world!")
		if output := writer.String(); strings.Contains(output, "Hello world!") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})
}

//nolint:dupl
func TestLogger_Warningf(t *testing.T) {
	t.Parallel()
	t.Run("unfiltered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelWarning, writer)
		logger.Warningf("Hello world!")
		if output := writer.String(); !strings.Contains(output, "Hello world!") || !strings.Contains(output, "warning") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})

	t.Run("filtered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelError, writer)
		logger.Warningf("Hello world!")
		if output := writer.String(); strings.Contains(output, "Hello world!") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})
}

func TestLogger_Errorf(t *testing.T) {
	t.Parallel()
	t.Run("unfiltered", func(t *testing.T) {
		t.Parallel()
		writer := log.NewBufferWriter()
		logger := log.NewLogger(log.LevelError, writer)
		logger.Errorf("Hello world!")
		if output := writer.String(); !strings.Contains(output, "Hello world!") || !strings.Contains(output, "error") {
			t.Fatalf("Incorrect log output: %s", output)
		}
	})
}

func TestLogger_WithLabel(t *testing.T) {
	t.Parallel()
	writer := log.NewBufferWriter()
	logger := log.NewLogger(log.LevelInfo, writer)
	logger2 := logger.WithLabel("source", "logger2")
	logger3 := logger2.WithLabel("foo", "bar")
	logger4 := logger3.WithLabel("foo", "baz")

	logger.Errorf("Hello world!")
	if output := writer.String(); strings.Contains(output, "source") || strings.Contains(output, "logger2") {
		t.Fatalf("Incorrect log output: %s", output)
	}

	logger2.Errorf("Hello world!")
	if output := writer.String(); !strings.Contains(output, "source") || !strings.Contains(output, "logger2") {
		t.Fatalf("Incorrect log output: %s", output)
	}
	if output := writer.String(); strings.Contains(output, "foo") || strings.Contains(output, "bar") {
		t.Fatalf("Incorrect log output: %s", output)
	}

	logger3.Errorf("Hello world!")
	if output := writer.String(); !strings.Contains(output, "foo") || !strings.Contains(output, "bar") {
		t.Fatalf("Incorrect log output: %s", output)
	}
	if output := writer.String(); strings.Contains(output, "baz") {
		t.Fatalf("Incorrect log output: %s", output)
	}

	logger4.Errorf("Hello world!")
	if output := writer.String(); !strings.Contains(output, "baz") {
		t.Fatalf("Incorrect log output: %s", output)
	}
}
