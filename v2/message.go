package log

import (
	"fmt"
	"time"
)

// Message is a single log message.
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	Level     Level     `json:"level"`
	Labels    Labels    `json:"labels"`
	Message   string    `json:"message"`
}

func (m Message) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", m.Timestamp.Format(time.RFC3339), m.Level, m.Labels, m.Message)
}
