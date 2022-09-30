package log

import (
    "fmt"
    goLog "log"
    "os"
    "time"
)

// Logger provides pluggable logging for Arcalot.
type Logger interface {
    Debugf(format string, args ...interface{})
    Infof(format string, args ...interface{})
    Warningf(format string, args ...interface{})
    Errorf(format string, args ...interface{})

    // WithLabel creates a child logger with this label attached.
    WithLabel(name string, value string) Logger
}

func New(config Config) Logger {
    if err := config.Validate(); err != nil {
        panic(err)
    }

    writer := newWriter(config)
    return NewLogger(config.Level, writer)
}

func newWriter(config Config) Writer {
    switch config.Destination {
    case DestinationStdout:
        out := config.Stdout
        if out == nil {
            out = os.Stdout
        }
        logger := goLog.New(out, "", 0)

        return NewGoLogWriter(logger)
    case DestinationTest:
        return NewTestWriter(config.T)
    default:
        panic(fmt.Errorf("invalid destination: %s", config.Destination))
    }
}

// NewLogger creates a new logger with the specified minimum level and target writer.
func NewLogger(minLevel Level, writer Writer) Logger {
    return &logger{
        minLevel,
        writer,
        map[string]string{},
    }
}

type logger struct {
    minLevel Level
    writer   Writer
    labels   Labels
}

func (l logger) Debugf(format string, args ...interface{}) {
    l.write(LevelDebug, format, args...)
}

func (l logger) Infof(format string, args ...interface{}) {
    l.write(LevelInfo, format, args...)
}

func (l logger) Warningf(format string, args ...interface{}) {
    l.write(LevelWarning, format, args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
    l.write(LevelError, format, args...)
}

func (l logger) write(level Level, message string, args ...interface{}) {
    if !l.minLevel.ShouldPrint(level) {
        return
    }
    if err := l.writer.Write(Message{
        Timestamp: time.Now(),
        Level:     level,
        Labels:    l.labels,
        Message:   fmt.Sprintf(message, args...),
    }); err != nil {
        panic(err)
    }
}

func (l logger) WithLabel(name string, value string) Logger {
    var newLabels Labels
    if _, contains := l.labels[name]; contains {
        newLabels = make(Labels, len(l.labels))
    } else {
        newLabels = make(Labels, len(l.labels)+1)
    }
    for k, v := range l.labels {
        newLabels[k] = v
    }
    newLabels[name] = value
    return &logger{
        minLevel: l.minLevel,
        writer:   l.writer,
        labels:   newLabels,
    }
}
