package smpsbuild

import "io"

type panicWriter struct {
	writer io.Writer
}

func newPanicWriter(w io.Writer) *panicWriter {
	pw := new(panicWriter)
	pw.writer = w
	return pw
}

func (pw *panicWriter) Write(p []byte) (n int, err error) {
	n, err = pw.writer.Write(p)
	if err != nil {
		panic(err)
	}

	return n, nil
}
