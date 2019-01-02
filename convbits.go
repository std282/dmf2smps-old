package main

import (
	"bytes"
	"io"
)

func releaseAccum(buf *bytes.Buffer, accum int) {
	if accum != 0 {
		for accum > 127 {
			buf.Write([]byte{127, 0xE7})
			accum -= 127
		}

		buf.WriteByte(byte(accum))
	}
}

func flushBuffer(buf *bytes.Buffer) []byte {
	if length := buf.Len(); length > 0 {
		slice := make([]byte, length)
		_, err := io.ReadFull(buf, slice)
		if err != nil {
			logger.Fatal("error: unable to flush non-empty buffer: ", err.Error())
		}

		return slice
	}

	return []byte{}
}
