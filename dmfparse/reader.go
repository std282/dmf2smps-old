package dmfparse

import (
	"io"
)

type safeReader struct {
	reader io.Reader
	err    error
}

func (r safeReader) Read(buf []byte) (n int, err error) {
	n, err = r.reader.Read(buf)
	if err != nil {
		panic(ReaderFailure{OutErr: err})
	}

	return
}

type ReaderFailure struct {
	OutErr error
}

func (err ReaderFailure) Error() string {
	return "failed to read DMF file: " + err.OutErr.Error()
}
