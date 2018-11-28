package dmfparse

import (
	"io"
)

func (song *Song) Parse(r io.Reader) {
	sr := safeReader{reader: r, err: nil}

}
