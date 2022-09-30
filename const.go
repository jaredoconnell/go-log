package log

import "fmt"

// Level is the logging level.
type Level string

const (
	// LevelDebug logs for debugging purposes. These logs are likely to contain excessive amounts of information.
	LevelDebug Level = "debug"
	// LevelInfo logs informational messages.
	LevelInfo Level = "info"
	// LevelWarning logs warning messages.
	LevelWarning Level = "warning"
	// LevelError logs error messages.
	LevelError Level = "error"
)

// ShouldPrint returns true if the a message at messageLevel should be printed if l is the current minimum level.
func (l Level) ShouldPrint(messageLevel Level) bool {
	switch l {
	case LevelError:
		return messageLevel == LevelError
	case LevelWarning:
		return messageLevel == LevelError ||
			messageLevel == LevelWarning
	case LevelInfo:
		return messageLevel == LevelError ||
			messageLevel == LevelWarning ||
			messageLevel == LevelInfo
	case LevelDebug:
		return messageLevel == LevelError ||
			messageLevel == LevelWarning ||
			messageLevel == LevelInfo ||
			messageLevel == LevelDebug
	default:
		panic(fmt.Errorf("invalid log level value: %s", l))
	}
}

func (l Level) Validate() error {
	switch l {
	case LevelError:
	case LevelWarning:
	case LevelInfo:
	case LevelDebug:
	default:
		return fmt.Errorf("invalid log level value: %s", l)
	}
	return nil
}

// Destination is the place a LogWriter writes to.
type Destination string

const (
	// DestinationStdout writes to the standard output.
	DestinationStdout Destination = "stdout"
	// DestinationTest writes the logs to the *testing.T facility.
	DestinationTest Destination = "test"
)

func (d Destination) Validate() error {
	switch d {
	case DestinationStdout:
	case DestinationTest:
	default:
		return fmt.Errorf("invalid destination value: %s", d)
	}
	return nil
}
