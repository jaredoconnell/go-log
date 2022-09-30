package log

import "testing"

// NewTestWriter returns a writer that logs via the Go test facility.
func NewTestWriter(t *testing.T) Writer {
	return newTargetWriter(t.Log)
}
