package dmfparse

import (
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
)

type panicReader struct {
	reader io.Reader
	eof    bool
}

func (pr *panicReader) handleError(err error) {
	switch err {
	case nil:
		return
	case io.EOF:
		pr.eof = true
	default:
		log.Fatal("dmfparse: error: reading error: ", err.Error())
	}
}

func (pr *panicReader) ReadAny(p interface{}) {
	err := binary.Read(pr.reader, binary.LittleEndian, p)
	pr.handleError(err)
}

func (pr *panicReader) Read8() byte {
	var b byte
	err := binary.Read(pr.reader, binary.LittleEndian, &b)
	pr.handleError(err)

	return b
}

func (pr *panicReader) Read8s() int8 {
	var b int8
	err := binary.Read(pr.reader, binary.LittleEndian, &b)
	pr.handleError(err)

	return b
}

func (pr *panicReader) Read16() int16 {
	var s int16
	err := binary.Read(pr.reader, binary.LittleEndian, &s)
	pr.handleError(err)

	return s
}

func (pr *panicReader) Read32() int32 {
	var i int32
	err := binary.Read(pr.reader, binary.LittleEndian, &i)
	pr.handleError(err)

	return i
}

func (pr *panicReader) Read32U() uint32 {
	var u uint32
	err := binary.Read(pr.reader, binary.LittleEndian, &u)
	pr.handleError(err)

	return u
}

func (pr *panicReader) ReadString(size int) string {
	buf := make([]byte, size)
	_, err := pr.reader.Read(buf)
	pr.handleError(err)
	return string(buf)
}

func (pr *panicReader) ReadPascalString() string {
	size := pr.Read8()
	return pr.ReadString(int(size))
}

func (pr *panicReader) Skip(bytes int) {
	io.CopyN(ioutil.Discard, pr.reader, int64(bytes))
}
