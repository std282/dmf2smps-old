package deflemask

import "io"

type panicReader struct {
	actualReader io.Reader
}

func newPanicReader(r io.Reader) *panicReader {
	return &panicReader{actualReader: r}
}

func (pr *panicReader) Read(p []byte) (int, error) {
	n, err := pr.actualReader.Read(p)
	if err != nil && err != io.EOF {
		internalPanic(err)
	}

	return n, err
}
