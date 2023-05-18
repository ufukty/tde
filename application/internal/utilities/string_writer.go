package utilities

import "bytes"

// StringWriter implements io.Writer interface
type StringWriter struct {
	buf    []byte
	Buffer *bytes.Buffer
}

func NewStringWriter() *StringWriter {
	sio := StringWriter{
		buf: []byte{},
	}
	sio.Buffer = bytes.NewBuffer(sio.buf)
	return &sio
}

func (sw *StringWriter) Write(p []byte) (n int, err error) {
	return sw.Buffer.Write(p)
}

func (sw *StringWriter) String() string {
	return sw.Buffer.String()
}
