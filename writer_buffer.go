package log

// BufferWriter provides a writer that records all log messages. This logger is most useful for tests.
type BufferWriter interface {
	Writer

	// String returns all recorded logs.
	String() string
}

func NewBufferWriter() BufferWriter {
	return &bufferWriter{}
}

type bufferWriter struct {
	buf []byte
}

func (b *bufferWriter) Write(message Message) error {
	b.buf = append(b.buf, []byte(message.String()+"\n")...)
	return nil
}

func (b *bufferWriter) Rotate() {

}

func (b *bufferWriter) Close() error {
	return nil
}

func (b *bufferWriter) String() string {
	return string(b.buf)
}
