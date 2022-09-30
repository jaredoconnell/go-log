package log

import (
    "fmt"
    "io"
    "testing"
)

// Config is the configuration for the logger.
type Config struct {
    // Level sets the minimum log level for the logger.
    Level Level `json:"level"`

    // Destination is the place the logger writes to.
    Destination Destination `json:"destination"`

    // T is the Go test for logging purposes.
    T *testing.T `json:"-" yaml:"-"`

    // Stdout is the standard output used by the DestinationStdout destination. If not set, os.Stdout is used.
    Stdout io.Writer `json:"-" yaml:"-"`
}

func (c Config) Validate() error {
    if err := c.Level.Validate(); err != nil {
        return fmt.Errorf("invalid level value (%w)", err)
    }
    if err := c.Destination.Validate(); err != nil {
        return fmt.Errorf("invalid destination value (%w)", err)
    }
    if c.Destination == DestinationTest && c.T == nil {
        return fmt.Errorf("no T value supplied with destination=test")
    }
    return nil
}
