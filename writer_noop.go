package log

// NewNOOPLogger returns a logger that does nothing.
func NewNOOPLogger() Writer {
	return &noopLogger{}
}

type noopLogger struct {
}

func (n noopLogger) Write(_ Message) error {
	return nil
}

func (n noopLogger) Rotate() {

}

func (n noopLogger) Close() error {
	return nil
}
