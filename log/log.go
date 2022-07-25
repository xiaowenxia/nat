package log

import (
	"io"
)

type alogWriter struct {
	w io.Writer
}

var (
	// AlogWriter Alog Writer
	AlogWriter      alogWriter
	writeBufferSize = 128 * 1024
)

func (a alogWriter) Printf(s string, i ...interface{}) {
	Sugar.Infow(s, i...)
}

func (a alogWriter) Write(p []byte) (int, error) {
	var sent int
	for len(p) > 0 {
		chunkSize := len(p)
		if chunkSize > writeBufferSize {
			chunkSize = writeBufferSize
		}

		Sugar.Info(string(p[:chunkSize]))
		sent += chunkSize
		p = p[chunkSize:]
	}
	return sent, nil
}

func (a alogWriter) ReadFrom(r io.Reader) (int64, error) {
	var nRead int64
	buf := make([]byte, writeBufferSize)

	var errRead, errSend error
	for errSend == nil && errRead != io.EOF {
		var n int
		n, errRead = r.Read(buf)
		nRead += int64(n)
		if errRead != nil && errRead != io.EOF {
			return nRead, errRead
		}

		if n > 0 {
			// errSend = s.sender(buf[n:])
		}
	}

	return nRead, errSend
}
