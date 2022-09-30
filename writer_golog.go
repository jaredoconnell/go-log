package log

import (
    "fmt"
    "log"
)

// NewGoLogWriter creates a log writer that writes to the Go log facility. The optional logger parameter can be
// used to pass one scoped logger, otherwise the global logger is used. If multiple loggers are passed the
// function will panic.
func NewGoLogWriter(logger ...*log.Logger) Writer {
    var l *log.Logger = nil
    if len(logger) == 1 {
        l = logger[0]
    } else if len(logger) == 0 {
        l = log.Default()
    } else {
        panic(fmt.Sprintf("Only one logger may be passed to NewGoLogger, %d were passed.", len(logger)))
    }
    return &targetWriter{
        l.Print,
    }
}
