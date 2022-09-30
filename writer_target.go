package log

// newTargetWriter returns a writer that logs via the common Target interface.
func newTargetWriter(logFunc func(...any)) Writer {
	return &targetWriter{
		logFunc: logFunc,
	}
}

type targetWriter struct {
	logFunc func(...any)
}

func (t targetWriter) Write(message Message) error {
	t.logFunc(message.String())
	return nil
}

func (t targetWriter) Rotate() {
}

func (t targetWriter) Close() error {
	return nil
}
