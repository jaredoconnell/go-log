package log

import "io"

// Writer is an abstraction over the possible logging targets.
type Writer interface {
	// Write writes a message to the output.
	Write(Message) error

	// Rotate gets called if logs need to be rotated.
	Rotate()

	io.Closer
}
