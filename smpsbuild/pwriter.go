package smpsbuild

/*
	This file describes panic writer.

	Panic writer is a simple writer, but it panics if returned error is not nil.
	Otherwise it is always successful. It is useful for a lot of writing operations,
	when you cannot check errors after every Write call.
*/

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
